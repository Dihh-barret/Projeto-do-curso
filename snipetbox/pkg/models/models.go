//colocar as coisas do banco de dados aqui trazendo para o programa e trazendo para uma variavel

package models

import ("time"
       "errors")

var ErrnoRecord = errors.New("models: no matching record Found")

type Snippet struct{
  ID int
  Title string
  Content string
  Created time.Time
  Expires time.Time
}