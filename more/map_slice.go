package more

import (
  "fmt"
  "testing"
)
var testTimeSlice=[]string{"aa","bb","cc","dd","ee","aa","zz"}

var testTimeMap = map[string]bool{"aa": true, "bb": true, "cc": true, "dd": true, "ee": true, "ff": true, "zz": true}

//以上为第一组查询测试数据


var testTimeSlice2=[] string{"aa","bb","cc","dd","ee","aa","aa","bb","cc","dd","ee","aa","aa","bb","cc","dd","ee","aa","aa","bb","cc","dd","ee","aa","i","j", "l", "m", "n", "o", "p", "q", "k", "x", "y", "z",
    "1", "2", "3", "4", "5", "6", "7", "8", "9", "10","zz"}

var testTimeMap2 = map[string]bool{"aa": true, "bb": true, "cc": true, "dd": true, "ee": true, "ff": true, "qq": true,"ww": true, "rr": true, "tt": true, "zz": true, "uu": true, "ii": true,"oo": true, "pp": true, "lk": true, "kl": true, "jk": true, "kj": true,"hl": true, "lh": true, "fg": true, "gfdd": true, "df": true, "fd": true,
    "i": true, "j": true, "l": true, "m": true, "n": true, "o": true, "p": true, "q": true, "k": true, "x": true, "y": true, "z": true,
    "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true}

//以上为第二组查询测试数据


func TestMoreSliceMap(t *T.testing) {
  TestSlice(testTimeSlice)
  TestMap(testTimeMap)
  
 TestSlice(testTimeSlice2)
 TestMap(testTimeMap2)

}


func TestSlice(a []string)  {
    now:=time.Now()

    for j:=0; j < 100000; j++{
        for _,v:=range a{
            if v=="zz"{
                break
            }
        }
    }
    finish:=time.Since(now)
fmt.Println(finish1)
}


func TestMap(a map[string]bool) {
    now:=time.Now()
    for j:=0; j < 100000; j++{
        if _, ok := a["zz"]; ok {
            continue
        }
    }
    finish2:=time.Since(now)
    fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
    fmt.Println(finish2)
    fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
}
