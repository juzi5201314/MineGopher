package nbt

type End struct {
	*NamedTag
}

type Byte struct {
	*NamedTag
}

type Short struct {
	*NamedTag
}

type Int struct {
	*NamedTag
}

type Long struct {
	*NamedTag
}

type Float struct {
	*NamedTag
}

type Double struct {
	*NamedTag
}

type String struct {
	*NamedTag
}

func NewEnd(name string) *End { return &End{NewNamedTag(name, EndTag, 0)} }

func NewByte(name string, value byte) *Byte { return &Byte{NewNamedTag(name, ByteTag, value)} }

func NewShort(name string, value int16) *Short { return &Short{NewNamedTag(name, ShortTag, value)} }

func NewInt(name string, value int32) *Int { return &Int{NewNamedTag(name, IntTag, value)} }

func NewLong(name string, value int64) *Long { return &Long{NewNamedTag(name, LongTag, value)} }

func NewFloat(name string, value float32) *Float { return &Float{NewNamedTag(name, FloatTag, value)} }

func NewDouble(name string, value float64) *Double {
	return &Double{NewNamedTag(name, DoubleTag, value)}
}

func NewString(name, value string) *String { return &String{NewNamedTag(name, StringTag, value)} }

func (tag *Byte) read(reader *Reader) { tag.value = reader.GetByte() }

func (tag *Short) read(reader *Reader) { tag.value = reader.GetShort() }

func (tag *Int) read(reader *Reader) { tag.value = reader.GetInt() }

func (tag *Long) read(reader *Reader) { tag.value = reader.GetLong() }

func (tag *Float) read(reader *Reader) { tag.value = reader.GetFloat() }

func (tag *Double) read(reader *Reader) { tag.value = reader.GetDouble() }

func (tag *String) read(reader *Reader) { tag.value = reader.GetString() }

func (tag *Byte) write(writer *Writer) { writer.PutByte(tag.value.(byte)) }

func (tag *Short) write(writer *Writer) { writer.PutShort(tag.value.(int16)) }

func (tag *Int) write(writer *Writer) { writer.PutInt(tag.value.(int32)) }

func (tag *Long) write(writer *Writer) { writer.PutLong(tag.value.(int64)) }

func (tag *Float) write(writer *Writer) { writer.PutFloat(tag.value.(float32)) }

func (tag *Double) write(writer *Writer) { writer.PutDouble(tag.value.(float64)) }

func (tag *String) write(writer *Writer) { writer.PutString(tag.value.(string)) }
