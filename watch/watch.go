package watch

import (
	_ "embed"
	"math"
	"time"
)

var (
	//go:embed foif.bmp
	foif []byte
	//go:embed zae.bmp
	zä []byte
	//go:embed viertel.bmp
	viertel []byte
	//go:embed zwaenzg.bmp
	zwänzg []byte
	//go:embed ab.bmp
	ab []byte
	//go:embed vor.bmp
	vor []byte
	//go:embed halbi.bmp
	halbi []byte
	//go:embed eis.bmp
	eis []byte
	//go:embed zwoi.bmp
	zwoi []byte
	//go:embed drue.bmp
	drü []byte
	//go:embed vieri.bmp
	vieri []byte
	//go:embed foifi.bmp
	foifi []byte
	//go:embed sechsi.bmp
	sechsi []byte
	//go:embed sebni.bmp
	sebni []byte
	//go:embed achti.bmp
	achti []byte
	//go:embed nueni.bmp
	nüni []byte
	//go:embed zaeni.bmp
	zäni []byte
	//go:embed oelfi.bmp
	ölfi []byte
	//go:embed zwoelfi.bmp
	zwölfi []byte
)

func Image(t time.Time) []byte {
	h := t.Hour()
	m := int(math.Round(float64(t.Minute())/5) * 5)
	switch m {
	case 0:
		return imageForHour(h)
	case 5:
		return merge(foif, ab, imageForHour(h))
	case 10:
		return merge(zä, ab, imageForHour(h))
	case 15:
		return merge(viertel, ab, imageForHour(h))
	case 20:
		return merge(zwänzg, ab, imageForHour(h))
	case 25:
		return merge(foif, vor, halbi, imageForHour(h+1))
	case 30:
		return merge(halbi, imageForHour(h+1))
	case 35:
		return merge(foif, ab, halbi, imageForHour(h+1))
	case 40:
		return merge(zwänzg, vor, imageForHour(h+1))
	case 45:
		return merge(viertel, vor, imageForHour(h+1))
	case 50:
		return merge(zä, vor, imageForHour(h+1))
	case 55:
		return merge(foif, vor, imageForHour(h+1))
	default:
		panic("unreachable")
	}
}

func imageForHour(h int) []byte {
	switch h % 12 {
	case 0:
		return zwölfi
	case 1:
		return eis
	case 2:
		return zwoi
	case 3:
		return drü
	case 4:
		return vieri
	case 5:
		return foifi
	case 6:
		return sechsi
	case 7:
		return sebni
	case 8:
		return achti
	case 9:
		return nüni
	case 10:
		return zäni
	case 11:
		return ölfi
	default:
		panic("unreachable")
	}
}

func merge(images ...[]byte) []byte {
	merged := images[0]
	for i := 1; i < len(images); i++ {
		for j := range images[i][62:] {
			merged[j] &= images[i][j]
		}
	}
	return merged
}
