package service

type Storage interface {
}

type GophKeeper struct {
	str Storage
}

func NewGophKeeper(str Storage) GophKeeper {
	return GophKeeper{
		str: str,
	}

}
