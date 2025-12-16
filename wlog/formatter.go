package wlog

type Formatter interface {
	Format(entry *Entry) error
}
