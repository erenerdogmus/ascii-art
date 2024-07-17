## ascii-art-justify

### Hedefler

İlk konudaki aynı [talimatları](../README.md) izlemelisiniz ancak hizalama değiştirilebilir.

```console
We
        will
                explain!
```

Çıktının hizalamasını değiştirmek için bir **bayrak** `--align=<type>` kullanmak mümkün olmalıdır; burada `type` şu şekilde olabilir:

- center
- left
- right
- justify

- Gösteriminizi terminal boyutuna göre uyarlamanız gerekir. Terminal penceresini küçültürseniz grafik gösterimi terminal boyutuna uyarlanmalıdır.
- Yalnızca terminal boyutuna uyan metin test edilecektir.
- Bayrak yukarıdakiyle tamamen aynı formatta olmalıdır; diğer formatlar aşağıdaki kullanım mesajını döndürmelidir:

```console
Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard
```

Uygulanan başka "ascii-art" isteğe bağlı projeler varsa, program diğer doğru biçimlendirilmiş "[SEÇENEK]" ve/veya "[BANNER]''ı kabul etmelidir.  
Ayrıca programın hâlâ tek bir "[STRING]" bağımsız değişkeniyle çalışabilmesi gerekir.


### Kullanım

Aşağıdaki ekrandaki çubukların terminal sınırları olduğunu varsayalım:

```console
|$ go run . --align=center "hello" standard                                                                                 |
|                                             _                _    _                                                       |
|                                            | |              | |  | |                                                      |
|                                            | |__      ___   | |  | |    ___                                               |
|                                            |  _ \    / _ \  | |  | |   / _ \                                              |
|                                            | | | |  |  __/  | |  | |  | (_) |                                             |
|                                            |_| |_|   \___|  |_|  |_|   \___/                                              |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=left "Hello There" standard                                                                             |
| _    _           _    _                 _______   _                                                                       |
|| |  | |         | |  | |               |__   __| | |                                                                      |
|| |__| |   ___   | |  | |    ___           | |    | |__      ___    _ __     ___                                           |
||  __  |  / _ \  | |  | |   / _ \          | |    |  _ \    / _ \  | '__|   / _ \                                          |
|| |  | | |  __/  | |  | |  | (_) |         | |    | | | |  |  __/  | |     |  __/                                          |
||_|  |_|  \___|  |_|  |_|   \___/          |_|    |_| |_|   \___|  |_|      \___|                                          |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=right "hello" shadow                                                                                    |
|                                                                                                                           |
|                                                                                          _|                _| _|          |
|                                                                                          _|_|_|     _|_|   _| _|   _|_|   |
|                                                                                          _|    _| _|_|_|_| _| _| _|    _| |
|                                                                                          _|    _| _|       _| _| _|    _| |
|                                                                                          _|    _|   _|_|_| _| _|   _|_|   |
|                                                                                                                           |
|                                                                                                                           |
|$ go run . --align=justify "how are you" shadow                                                                            |
|                                                                                                                           |
|_|                                                                                                                         |
|_|_|_|     _|_|   _|      _|      _|                  _|_|_| _|  _|_|   _|_|                    _|    _|   _|_|   _|    _| |
|_|    _| _|    _| _|      _|      _|                _|    _| _|_|     _|_|_|_|                  _|    _| _|    _| _|    _| |
|_|    _| _|    _|   _|  _|  _|  _|                  _|    _| _|       _|                        _|    _| _|    _| _|    _| |
|_|    _|   _|_|       _|      _|                      _|_|_| _|         _|_|_|                    _|_|_|   _|_|     _|_|_| |
|                                                                                                      _|                   |
|                                                                                                  _|_|                     |
|$                                                                                                                          |
```

Bu proje aşağıdakileri öğrenmenize yardımcı olacaktır:

- Go dosya sistemi(**fs**) API'si
- Veri manipülasyonu
- Terminal ekranı