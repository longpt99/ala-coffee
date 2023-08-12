package product

import (
	"context"
	"ecommerce/configs/database"
	"fmt"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	InsertOne(params interface{}) (*string, error)
	List() (*[]Product, error)
	DetailByID(id string) (*Product, error)
	Delete(id string) error
}

type QueryParams struct {
	Columns   []string
	TableName string
	Where     string
	OrderBy   string
	Limit     int
	Offset    int
	Args      []interface{}
}

type FieldData struct {
	Key   string
	Value interface{}
}

type repo struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func InitRepository(store *database.PostgresConfig) Repository {
	return &repo{
		db:  store.DB,
		ctx: store.Ctx,
	}
}

func (r *repo) List() (*[]Product, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT * 
		FROM products 
		WHERE deleted_at IS NULL
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[Product])

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *repo) DetailByID(id string) (*Product, error) {
	rows, err := r.db.Query(context.Background(), `
		SELECT * FROM products WHERE id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	data, err := pgx.CollectRows(rows, pgx.RowToStructByName[Product])

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, nil
	}

	return &data[0], err
}

func (r *repo) Delete(id string) error {
	_, err := r.db.Exec(context.Background(), `
		UPDATE products SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1
	`, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *repo) InsertOne(params interface{}) (*string, error) {
	var id string

	data := r.getStructFields(params)
	dataLength := len(data)
	args := make([]interface{}, dataLength)
	valueStrings := make([]string, dataLength)
	keys := make([]string, dataLength)

	for i, d := range data {
		valueStrings = append(valueStrings, fmt.Sprintf("$%d", i+1))
		keys = append(keys, d.Key)
		args = append(args, d.Value)
	}

	query := fmt.Sprintf(`
	INSERT INTO users(%s)
	VALUES (%s)
	RETURNING id;
	`, strings.Join(keys, ","), strings.Join(valueStrings, ","))

	err := r.db.QueryRow(context.Background(), query, args...).Scan(&id)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &id, nil
}

func (*repo) generateQuery(params *QueryParams) (string, []interface{}) {
	var sb strings.Builder

	sb.WriteString("SELECT ")

	if len(params.Columns) == 0 {
		sb.WriteString("*")
	} else {
		sb.WriteString(strings.Join(params.Columns, ", "))
	}

	sb.WriteString(fmt.Sprintf(" FROM %s", params.TableName))

	if params.Where != "" {
		sb.WriteString(fmt.Sprintf(" WHERE %s", params.Where))
	}

	if params.OrderBy != "" {
		sb.WriteString(fmt.Sprintf(" ORDER BY %s", params.OrderBy))
	}

	if params.Limit > 0 {
		sb.WriteString(fmt.Sprintf(" LIMIT %d", params.Limit))
	}

	if params.Offset > 0 {
		sb.WriteString(fmt.Sprintf(" OFFSET %d", params.Offset))
	}

	query := sb.String()

	return query, params.Args
}

func (r *repo) strutForScan(u interface{}) []interface{} {
	val := reflect.ValueOf(u).Elem()
	v := make([]interface{}, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		v[i] = valueField.Addr().Interface()
	}

	return v
}

func (r *repo) getStructFields(data interface{}) []FieldData {
	var fields []FieldData

	value := reflect.ValueOf(data)
	if value.Kind() != reflect.Struct {
		fmt.Println("Not a struct.")
		return nil
	}

	typeOf := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldName := typeOf.Field(i).Name

		if field.Kind() == reflect.Struct {
			embeddedFields := r.getStructFields(field.Interface())
			fields = append(fields, embeddedFields...)
		} else {
			fields = append(fields, FieldData{Key: strings.ToLower(fieldName), Value: field.Interface()})
		}
	}

	return fields
}
