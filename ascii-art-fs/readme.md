## ascii-art-fs

### Hedefler

İlk konuda olduğu gibi aynı [talimatları](../README.md) takip etmelisiniz, ancak ikinci argüman şablonun adı olmalıdır. Bazı şablonların okunmasının zor olabileceğini biliyorum, sadece bunu takıntı haline getirmeyin.

### Talimatlar

- Projeniz **Go** dilinde yazılmış olmalıdır.
- Kod [**iyi uygulamalara**](../../good-practices/README.md) uymalıdır.
- Birim testi](https://go.dev/doc/tutorial/add-a-test) için **test dosyalarının** olması önerilir.
- Bannerlar** hakkında her şeyi [burada](../) görebilirsiniz.
- Kullanım bu formata uymalıdır `go run . [STRING] [BANNER]` biçimine uymalıdır, diğer biçimler aşağıdaki kullanım mesajını döndürmelidir:

```console
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```

Uygulanan başka `ascii-art` isteğe bağlı projeler varsa, program doğru biçimlendirilmiş diğer `[OPTION]` ve/veya `[BANNER]` argümanlarını kabul etmelidir.  
Ek olarak, program hala tek bir `[STRING]` argümanı ile çalışabilmelidir.

### Kullanım

```console
student$ go run . "hello" standard | cat -e
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
student$ go run . "Hello There!" shadow | cat -e
                                                                                         $
_|    _|          _| _|                _|_|_|_|_| _|                                  _| $
_|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
_|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
_|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
_|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                         $
                                                                                         $
student$ go run . "Hello There!" thinkertoy | cat -e
                                                $
o  o     o o           o-O-o o                o $
|  |     | |             |   |                | $
O--O o-o | | o-o         |   O--o o-o o-o o-o o $
|  | |-' | | | |         |   |  | |-' |   |-'   $
o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                $
                                                $
```

Bu proje hakkında bilgi edinmenize yardımcı olacaktır:

- Go dosya sistemi(**fs**) API'si
- Veri manipülasyonu