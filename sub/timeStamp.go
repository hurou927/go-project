package timeStamp
import (
    "time"
    "fmt"
)
type TimeStamp struct{
    Length int
    Times [] time.Time
    Index int
}

func NewTimeStamp(l int) (*TimeStamp){
    var ts TimeStamp
    ts.Length = l
    ts.Times = make([]time.Time, l, l)
    ts.Index = 0
    return &ts
}

func (this *TimeStamp) Stamp () (){
    this.Times[this.Index] = time.Now()
    this.Index = this.Index + 1 ;
}
func (this *TimeStamp) Interval (p int, n int) (float64){
    return (this.Times[n].Sub(this.Times[p])).Seconds()*1000
}

func (this *TimeStamp) Print () (){
    for i:= 0; i < this.Index-1; i++ {
        fmt.Printf("%f,ms\n",this.Interval(i,i+1));
    }
}

func (this *TimeStamp) Print_str ( strs  []string ) (){

    var lengthOfstrs int = len(strs)
    for i:= 0; i < this.Index-1; i++ {
        if i < lengthOfstrs {
            fmt.Printf("%s,",strs[i])
        }else{
            fmt.Printf(",")
        }
        fmt.Printf("%f,ms,",this.Interval(i,i+1))
    }
    fmt.Printf("\n");
}

func (this *TimeStamp) Print_oneLine () (){
    for i:= 0; i < this.Index-1; i++ {
        fmt.Printf("%f,ms\n",this.Interval(i,i+1));
    }
}

func (this *TimeStamp) Print_str_oneLine ( strs  []string ) (){

    var lengthOfstrs int = len(strs)
    for i:= 0; i < this.Index-1; i++ {
        if i < lengthOfstrs {
            fmt.Printf("%s,",strs[i])
        }else{
            fmt.Printf(",")
        }
        fmt.Printf("%f,ms,",this.Interval(i,i+1))
    }
    fmt.Printf("\n");
}
//
// func()
