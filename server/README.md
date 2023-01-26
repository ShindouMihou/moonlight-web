<div align="center">
    once the darkness floods the skies, the moon will shine once again
</div>

##
### 〰️ Midnight

Midnight is the official name for the server client of Moonlight, it is the brains of Moonlight and handles communication 
between the database and the client from storing chapters, books and all other mess. It is written completely in Golang with 
little pain.

### 🔬 Contributing

You can contribute to the project by adding your own features, improvements or fixes to the server here. To run the project, you 
can simply run the following command:
```shell
go build app.go
```

### 👀 Key Notes
For implementing clients such as the Sunshine client, please keep the following keynotes since they can be destructive:
- `PATCH /chapters/:id` will set the contents to an empty string if no contents is included.
- All routes do not accept empty strings due to some technical issues, and will be intended to keep that way.