# mcore

**MCore** is a minimal storage core written in Go.

It provides simple data storage and retrieval using a local database.

---

### Features

- Store data units
- Retrieve data by ID
- Uses local storage (bbolt)

### Data Model

Each stored object contains:

- `ID` — unique identifier  
- `Type` — type of data (note, link, secret, etc.)  
- `Content` — main data  
- `Tags` — optional labels  
- `CreatedAt` — timestamp

### Example

engine, _ := mcore.NewEngine("test.db")

engine.Save(mcore.DataUnit{
    ID: "1",
    Type: "note",
    Content: "Hello MCore!",
})

data, _ := engine.Get("1")
fmt.Println(data.Content).


### Installation

```bash
go get github.com/maruf-mlon/mcore
