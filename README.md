# LeapYearGo

```golang
func isLeapYear(year int) bool {
	return ((year%400 == 0) || ((year%100 != 0) && (year%4 == 0)))
}
```