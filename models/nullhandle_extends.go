package models
////*
//Article struct
//type Article struct {
//	ID      int        `json:"id"`
//	Title   string     `json:"title"`
//	PubDate NullTime   `json:"pub_date"`
//	Body    NullString `json:"body"`
//	User    NullInt64  `json:"user"`
//}
//
//// NullInt64 is an alias for sql.NullInt64 data type
//type NullInt64 struct {
//	sql.NullInt64
//}
//
//// MarshalJSON for NullInt64
//func (ni *NullInt64) MarshalJSON() ([]byte, error) {
//	if !ni.Valid {
//		return []byte("null"), nil
//	}
//	return json.Marshal(ni.Int64)
//}
//
//// UnmarshalJSON for NullInt64
//// func (ni *NullInt64) UnmarshalJSON(b []byte) error {
////  err := json.Unmarshal(b, &ni.Int64)
////  ni.Valid = (err == nil)
////  return err
//// }
//
//// NullBool is an alias for sql.NullBool data type
//type NullBool struct {
//	sql.NullBool
//}
//
//// MarshalJSON for NullBool
//func (nb *NullBool) MarshalJSON() ([]byte, error) {
//	if !nb.Valid {
//		return []byte("null"), nil
//	}
//	return json.Marshal(nb.Bool)
//}
//
//// UnmarshalJSON for NullBool
//// func (nb *NullBool) UnmarshalJSON(b []byte) error {
////  err := json.Unmarshal(b, &nb.Bool)
////  nb.Valid = (err == nil)
////  return err
//// }
//
//// NullFloat64 is an alias for sql.NullFloat64 data type
//type NullFloat64 struct {
//	sql.NullFloat64
//}
//
//// MarshalJSON for NullFloat64
//func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
//	if !nf.Valid {
//		return []byte("null"), nil
//	}
//	return json.Marshal(nf.Float64)
//}
//
//// UnmarshalJSON for NullFloat64
//// func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
////  err := json.Unmarshal(b, &nf.Float64)
////  nf.Valid = (err == nil)
////  return err
//// }
//
//// NullString is an alias for sql.NullString data type
//type NullString struct {
//	sql.NullString
//}
//
//// MarshalJSON for NullString
//func (ns *NullString) MarshalJSON() ([]byte, error) {
//	if !ns.Valid {
//		return []byte("null"), nil
//	}
//	return json.Marshal(ns.String)
//}
//
//// UnmarshalJSON for NullString
//// func (ns *NullString) UnmarshalJSON(b []byte) error {
////  err := json.Unmarshal(b, &ns.String)
////  ns.Valid = (err == nil)
////  return err
//// }
//
//// NullTime is an alias for mysql.NullTime data type
//type NullTime struct {
//	mysql.NullTime
//}
//
//// MarshalJSON for NullTime
//func (nt *NullTime) MarshalJSON() ([]byte, error) {
//	if !nt.Valid {
//		return []byte("null"), nil
//	}
//	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
//	return []byte(val), nil
//}
//
//// UnmarshalJSON for NullTime
//// func (nt *NullTime) UnmarshalJSON(b []byte) error {
////  err := json.Unmarshal(b, &nt.Time)
////  nt.Valid = (err == nil)
////  return err
//// }
//
//// MAIN program starts here
//func main() {
//	log.Println("connectiong to database...")
//	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/test?charset=utf8")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	// read articles
//	rows, err := db.Query("SELECT * FROM Article")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer rows.Close()
//
//	// loop over articles
//	for rows.Next() {
//		var a Article
//		if err = rows.Scan(&a.ID, &a.Title, &a.PubDate, &a.Body, &a.User); err != nil {
//			log.Fatal(err)
//		}
//
//		log.Printf("article instance := %#v\n", a)
//		articleJSON, err := json.Marshal(&a)
//		if err != nil {
//			log.Fatal(err)
//		} else {
//			log.Printf("json marshal := %s\n\n", articleJSON)
//		}
//	}
//
//	err = rows.Err()
//	if err != nil {
//		log.Fatal(err)
//	}
//}*/
