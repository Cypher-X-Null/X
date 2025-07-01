package cypher

// این بخش نیاز به طراحی دقیق‌تر دارد و بسته به ساختار پیام‌ها باید 
//struct و متدهای مناسب تعریف شود.
// به عنوان نمونه:
type TLQueryHeader struct {
	QID     int64
	ActorID int64
	Flags   int
	Op      int
	RealOp  int
	RefCnt  int
}

func ParseTLQueryHeader(data []byte) (*TLQueryHeader, error) {
	// پیاده‌سازی پارسینگ بر اساس پروتکل
	// اینو خودت هندش کن آقای ایکس عزیز 
	return nil, nil
}