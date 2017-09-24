package distconv

//MToFt konwertuje dystans w metrach na stopy
func MToFt(m Meter) Foot { return Foot(m * 3.2808) }

//FtToM konwertuje dystans w stopach na metry
func FtToM(f Foot) Meter { return Meter(f / 3.2808) }
