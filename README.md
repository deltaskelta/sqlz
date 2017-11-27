# SQLZ

This package contains a few helper functions that are useful in conjuntion with [sqlx](github.com/jmoiron/sqlx).

# Examples:

## Generate Update SQL based on a struct

```go
    type Foo struct {
        ID int        `db:"id"`
        Field1 string `db:"field_1"`
        FIeld2 int    `db:"field_2"`
    }


func Update() error {
    foo := Foo{
        ID: 1,
        Field1: "foo",
    }

    sql, err := sqlz.UpdateSQL(foo, "footable")
    if err != nil {
        return err
    }

    // sql: UPDATE footable SET field_1 = :field_1 WHERE id = 1

    // db is a db created by sqlx.Connect
    result, err := db.NamedExec(sql, &foo)
    if err != nil {
        return err
    }
    
	aff, err := result.RowsAffected()
	if aff == 0 {
		return errors.New("0 rows afftected by update")
	}
	if err != nil {
		return errors.Wrap(err, "UpdateUser")
	}
}
```

## Use Slice Type To Store Slices in a Database

Slices need to be marshaled and unmarshaled because most databases, or the drivers in go
do not support storing slices in the actual db

```go
type Foo struct {
    Field1 sqlz.StringSlice `db:"field_1"`
}
```

because the `sqlz.StringSlice` implements the `driver.Valuer` interfaces, the rest will be
taken care of as it gets put into and pulled from the database by the driver.
