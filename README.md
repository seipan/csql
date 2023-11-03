<div align="center">

![Last commit](https://img.shields.io/github/last-commit/seipan/csql?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/seipan/csql?style=flat-square)
![Issues](https://img.shields.io/github/issues/seipan/csql?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/seipan/csql?style=flat-square)
[![go](https://github.com/seipan/csql/actions/workflows/go.yml/badge.svg)](https://github.com/seipan/csql/actions/workflows/go.yml)
![codecov](https://codecov.io/gh/seipan/csql/graph/badge.svg?token=6TCAKD8LY7)

<img src="./logo/csqllogo.png" alt="eyecatch" height="200">

# csql

⭐ CLI tool to insert CSV data into a specified database.  ⭐

<br>
<br>


</div>

## Install
```
go install github.com/seipan/csql@latest
```

## Usage
command option
```
Usage:
  csql [flags]

Flags:
  -c, --check         check csv format
  -d, --dns string    DNS for Connecting Database
  -h, --help          help for csql
  -p, --path string   FilePath for Parsing CSVFile
  -q, --query         output query
  -t, --type string   Database Type
```

### ```--check``` option
if success patern
```
csql --check --path=./testdata/csv/test01.csv --type=mysql --dsn=hogehoge

             ___________ ____    __ 
            / ____/ ___// __ \  / / 
           / /    \__ \/ / / / / /  
          / /___ ___/ / /_/ / / /___
          \____//____/\___\_\/_____/
                                                                          
                                                                   

csv format is correct
```
failed pattern
```
csql --check --path=./testdata/csv/test02.csv --type=mysql --dsn=hogehoge

             ___________ ____    __ 
            / ____/ ___// __ \  / / 
           / /    \__ \/ / / / / /  
          / /___ ___/ / /_/ / / /___
          \____//____/\___\_\/_____/
                                                                          
                                                                   

csv format is incorrect : table name is empty
exit status 1
```
### ```--dns``` option
### ```--path``` option
### ```--query``` option
### ```--type``` option


 ## License
Code licensed under 
[the MIT License](https://github.com/seipan/csql/blob/main/LICENSE).


## Author
[seipan](https://github.com/seipan).
