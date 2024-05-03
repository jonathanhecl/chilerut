# ChileRut on Golang
Golang package that provides common functionality regarding Chilean RUT

# Installation
``` 
$ go get github.com/jonathanhecl/chilerut
```

# Usage

```
import "github.com/jonathanhecl/chilerut"
```

## Validate RUT
```
chilerut.Valid("12345678-9")
# returns False
chilerut.Valid("6265837-1")
# returns True

# It works with the following formats

chilerut.Valid("98685030")
# returns True
chilerut.Valid("9868503-0")
# returns True
chilerut.Valid("9.868.503-0")
# returns True
chilerut.Valid("12-667-869-K")
# returns True
chilerut.Valid("12*667*869-k")
# returns True
```

## Get verification digit

```
chilerut.VerificationDigit("9868503")
# returns "0"
chilerut.VerificationDigit("12667869")
# returns "K"
chilerut.VerificationDigit("12667869")
# returns "K"
```

## Format RUT

```
chilerut.Format("12.667.869.k")
# returns "12667869-K"
chilerut.Format("12-667-869-k")
# returns "12667869-K"
chilerut.Format("12*667*869*K")
# returns "12667869-K"
chilerut.Format("12 667 869 k")
# returns "12667869-K"
chilerut.Format("12667869k")
# returns "12667869-K"
chilerut.Format("   000012667869k   ")
# returns "12667869-K"
```

## Compare RUT

```
chilerut.Compare("12667869k", "12.667.869-K")
# return True
chilerut.Compare("12667869-k", "12.667.869-K")
# return True
chilerut.Compare("12.667.861-K", "12.667.869-K")
# return False
```