## ascii-art-color

### Hedefler

İlk konuda olduğu gibi aynı [talimatları](../README.md) izlemelisiniz, ancak bu sefer renklerle.

Çıktı, `--color=<color> <substring to be colored>` **bayrağını** kullanarak renkleri manipüle etmelidir; burada `--color` bayraktır ve `<color>` kullanıcı tarafından istenen renktir ve `<substring to be colored>` renklendirilmek üzere seçebileceğiniz alt dizedir. Bu renkler farklı notasyonlar (`RGB`, `hsl`, `ANSI`... gibi renk kodu sistemleri) kullanılarak elde edilebilir, hangisini kullanmak istediğinizi seçmek size kalmıştır.

- Renklendirilecek alt dize tek bir harf veya daha fazla olabilir
- Alt dize belirtilmezse, tüm `dize` renklendirilmelidir.
- Bayrak tam olarak yukarıdaki biçimle aynı olmalıdır, diğer biçimler aşağıdaki kullanım mesajını döndürmelidir:

```console
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```

### Kullanım
```shell
$ go run . --color=red kit "a king kitten have kit"
```

Yukarıdaki örnek için, `kitten` kelimesindeki `kit` alt dizesi ve sondaki `kit` kelimesi renklendirilmelidir.

Uygulanan başka `ascii-art` opsiyonel projeleri varsa, program diğer doğru biçimlendirilmiş `[OPTION]` ve/veya `[BANNER]`ları kabul etmelidir.
Ayrıca, program yine de tek bir `[STRING]` argümanıyla çalışabilmelidir.

Bu proje hakkında bilgi edinmenize yardımcı olacaktır:

- Go dosya sistemi(**fs**) API'si
- Renk dönüştürücüler
- Veri manipülasyonu
- Terminal ekranı