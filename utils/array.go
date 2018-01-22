package utils

import(
    "errors"
    "sync"
)


type Array []interface{}

type ItemArray struct {
    items Array
    sync.Mutex
}


func (this *ItemArray) Add(elem interface{}) {
    this.Lock()
    this.items.Add(elem)
    this.Unlock()
}


func (this *ItemArray) AddAt(elem interface{}, index int) error{
    this.Lock()
    err := this.items.AddAt(elem, index)
    this.Unlock()
    return err
}


func (this *ItemArray) RemoveAt(index int) (interface{} , error){
    this.Lock()
    result,err := this.items.RemoveAt(index)
    this.Unlock()
    return result,err
}

func (this *ItemArray) Remove(elem interface{}) error{
    this.Lock()
    err := this.items.Remove(elem)
    this.Unlock()
    return err
}


func (this *ItemArray) Get(index int) (interface{},error) {
    return this.items.Get(index)
}

func (this *ItemArray) Size() int {
    return this.items.Size()
}

func NewArray() *ItemArray{
    return &ItemArray{items:make(Array,0)}
}

func (this *Array) Add(elem interface{}){
    *this = append(*this, elem)
}

func (this *Array) AddAt(elem interface{}, index int) error{
    count := this.Size()
    if(index < 0 || index >= count) {
       return errors.New("index not in array bound")
    }
    *this = append((*this)[0:index],elem,(*this)[index+1:count])
    return nil
}


func (this *Array) RemoveAt(index int) (interface{} , error){
    count := this.Size()
    if(index < 0 || index >= count) {
       return nil,errors.New("index not in array bound")
    }
    elem := (*this)[index]
    *this = append((*this)[:index] , (*this)[index+1:count]...)
    return elem,nil
}

func (this *Array) Remove(elem interface{}) error{
    count := this.Size()
    found := false
    for i,v := range *this {
        if v == elem {
            if i == count -1 {
                *this = (*this)[:i]
            } else {
                *this = append((*this)[:i] , (*this)[i+1:count]...)
            }
            found = true
            break
        }
    }
    if !found {
        return errors.New("not such element")
    }
    return nil
}

func (this *Array) Get(index int) (interface{},error) {
    count := this.Size()
    if(index < 0 || index >= count) {
       return nil,errors.New("index not in array bound") 
    }
    return (*this)[index], nil
}

func (this *Array) Size() int {
    return len(*this)
}

