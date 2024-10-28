package main

func main() {

}

func SetAlarm(employed, vacation bool) bool {
	if employed == true && vacation == true {
	  return false
	}
	if employed == false && vacation == true {
		return false
	}
	if employed == false && vacation == false {
		return false
	}
	if employed == true && vacation == false {
		return true
	}
	return false
  }


  /* 
setAlarm(true, true) -> false
setAlarm(false, true) -> false
setAlarm(false, false) -> false
setAlarm(true, false) -> true