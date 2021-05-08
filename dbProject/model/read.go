package model

import "workspace/dbProject/view"

func ReadAll() ([]view.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}
	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil

}
func ReadByQuery(name string) ([]view.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?", name)
	if err != nil {
		return nil, err
	}
	todos := []view.PostRequest{}
	for rows.Next() {
		data := view.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}
	return todos, nil

}
