package structs

type SortItem struct {
	Num  int
	Line string
}

type DataChanItem struct {
	SortItem
	Err error
}
