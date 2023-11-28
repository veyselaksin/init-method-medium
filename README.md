# Go Dilinde Başlangıç Ritüeli: Init Fonksiyonuyla Projenizin Temelini Atın

Go dilini öğrenmeye veya bir Go projesi üzerinde çalışmaya yeni başladıysanız, karşınıza çıkan ilk adımlardan biri `init` fonksiyonu olabilir. Bu fonksiyon, projenizin temelini atmanızı sağlar ve projenizin düzgün çalışması aynı zamanda bağımlılıkların yönetilmesi için önemli bir rol oynar. Bu yazıda, `init` fonksiyonunun ne olduğunu, nasıl kullanıldığını ve `init` fonksiyonu ile bir Go projesinin nasıl başlatılabileceğini detaylı bir şekilde ele alacağız.

## İçerikler

1. **Go Dilinde `init` Fonksiyonu Nedir?**
2. `**go mod init` Komutu ve Modül Oluşturma**
3. `**init` Fonksiyonunun Kullanımı ve Rolü**
4. `**init` Fonksiyonunda Değişkenlere İlk Değer Atama**
5. `**init` Fonksiyonu Kullanım Örneği: Birden Fazla Dosyada `init` Fonksiyonu Kullanımı ve Çalışma Sırası**

### 1. Go Dilinde `init` Fonksiyonu Nedir?

Go dilinde `init` fonksiyonu, bir paketin başlatılma sürecinde özel bir işlevi ifade eder. Herhangi bir paketin içinde `init` fonksiyonu tanımlanabilir, ve bu fonksiyon go tarafından otomatik olarak çalıştırılır. `init` fonksiyonu, genellikle paketin başlatılması sırasında yapılması gereken hazırlıkları gerçekleştirmek için kullanılır.

`init` fonksiyonları, bir paketin içindeki diğer kodlardan önce çalıştırılır ve bir paketin yalnızca bir `init` fonksiyonu olabilir. Bu fonksiyonlar, paketin içinde tanımlandıkları sırayla çalıştırılır.

Örnek olarak, bir paket içinde şu şekilde bir `init` fonksiyonu tanımlanabilir:

```go
package mypackage

import "fmt"

func init() {
    fmt.Println("Bu init fonksiyonu çalıştırıldı.")
}

```

Bu `init` fonksiyonu, `mypackage` paketi başlatıldığında otomatik olarak çalışacaktır. `init` fonksiyonları genellikle paketin başlatılması için gerekli olan ön hazırlıkları yapmak veya belirli bir sırayla işlemleri gerçekleştirmek için kullanılır.

### 2. `**go mod init` Komutu ve Modül Oluşturma**

Bir Go projesi oluşturulurken, genellikle ilk adım `go mod init` komutunu kullanmaktır. Bu komut, projeniz için bir modül oluşturur. Bir modül, belirli bir yolu olan ve Go kodunuzun bir versiyonunu içeren bir koleksiyondur. `go mod init` komutunu şu şekilde kullanabilirsiniz:

```go
go mod init example.com/myproject

```

Bu komut, `example.com/myproject` adında yeni bir modül oluşturur. Ayrıca, modülün adını ve kullanılacak Go sürümünü belirten bir `go.mod` dosyası oluşturur. Bu dosya, projenizin kök dizininde yer almalıdır. `go.mod` dosyası, projenizin bağımlılıklarını da izler. Bir bağımlılık eklediğinizde, `go.mod` dosyası otomatik olarak güncellenir. Bu sayede oluşturulan proje yapısal olarak daha yönetilebilir olur.

### 3. `**init` Fonksiyonunun Kullanımı ve Rolü**

Bir `init` fonksiyonunun kullanımı, genellikle bir programın başında, belli başlı ayarlamaları ve kontrolleri otomatik olarak yapmak için kullanılır. Ancak, `init` fonksiyonunun gereksiz yere veya yanlış bir şekilde kullanılması, programın karmaşıklığını artırabilir ve hatalara yol açabilir. Bu nedenle, `init` fonksiyonunun ne zaman ve nasıl kullanılacağına dikkatli bir şekilde karar vermek önemlidir.

Go dilinde, `init` fonksiyonu özel bir fonksiyondur ve her paketin içinde sadece bir kez çağrılır. Bu fonksiyon, paketin içindeki diğer tüm fonksiyonlardan önce otomatik olarak çağrılır. En sık, `init` fonksiyonu kullanım senaryosu ise değişkenlerin ilk değerlerini ayarlamak veya belirli bir durumun mevcudiyetini kontrol etmek için kullanılır. Kullanılması gerektiğinde, `init` fonksiyonu belirli bir dosyanın içinde veya paket seviyesinde tanımlanabilir ve herhangi bir giriş parametresi veya dönüş değeri olmaz.

`init` fonksiyonunun oluşturulması ve kullanılması oldukça basittir. Öncelikle, fonksiyonu tanımlamak için aşağıdaki gibi bir kod yazmanız gerekmektedir:

```go
func init() {
    // fonksiyonun yapacağı işlemler
}

```

Bu kod parçası, `init` fonksiyonunun biçimidir ve içerisine fonksiyonun yapacağı işlemler eklenmelidir. `init` fonksiyonu, giriş parametresi veya dönüş değeri olmadığı için, yapacağı işlemler genellikle değişkenlerin ilk değerlerini ayarlamak veya belirli bir durumun mevcudiyetini kontrol etmek olacaktır.

### 4. `**init` Fonksiyonunda Değişkenlere İlk Değer Atama**

Örneğin, bir `init` fonksiyonunun bir değişkene ilk değerini ayarlama işlemi aşağıdaki gibi olabilir:

```go
var number int

func init() {
    number = 10
}

```

Bu kod parçası, `number` adlı bir tamsayı değişkeni oluşturur ve `init` fonksiyonu içerisinde bu değişkenin ilk değerini 10 olarak ayarlar. Bu sayede, `number` değişkeni programın diğer yerlerinde kullanıldığında, ilk değeri 10 olacaktır.

`init` fonksiyonu, bir paketin içerisinde birden fazla dosya bulunuyorsa ve her bir dosyanın bir `init` fonksiyonu varsa, bu `init` fonksiyonları dosyaların yüklenme sırasına göre çağrılır. Yani, dosyaların yüklenme sırasına göre `init` fonksiyonları çalışır.

### `**init` Fonksiyonu Kullanım Örneği: Birden Fazla Dosyada `init` Fonksiyonu Kullanımı ve Çalışma Sırası**

`init` fonksiyonunun özellikleri ve basit bir kullanım örneği hakkında bilgi verildi. Ancak, daha karmaşık bir proje senaryosunda `init` fonksiyonunun nasıl kullanılabileceğini daha iyi anlamak için, `init` fonksiyonunun birden fazla dosya içinde kullanıldığı bir durumu düşünelim.

Kodumuzun dosya yapısı aşağıdaki gibidir:

```jsx
.
└── example.com/myproject/
    └── pkg/
        ├── calculate/
        │   └── calculate.go
        ├── greeting/
        │   └── greeting.go
        ├── helloworld/
        │   └── helloworld.go
        ├── go.mod
        └── main.go
```

Örneğin, `main.go`, `calculate.go`, `greeting.go` ve `helloworld.go` adında üç dosyanız olduğunu hayal edelim. Her bir dosyada, belirli işlemler gerçekleştiren bir `init` fonksiyonu tanımlanmıştır. Ayrıca her bir dosyada bir iş icra eden fonskiyonlarımız yer almaktadır.

```go
// calculate.go
package calculate

func init() {
	println("First init in calculate package")
}

// Add function adds two integers and returns the result
func Add(a, b int) int {
	return a + b
}
```

```go
// greeting.go
package greeting

func init() {
	println("First init in greeting package")
}

// SayHello function prints a greeting message
func SayHello() {
	println("Hello from greeting package")
}
```

```go
//helloworld.go
package helloworld

import "fmt"

var (
	name string
)

func init() {
	name = "Veysel"
	println("First init in helloworld package")
}

func SayHello() {
	fmt.Printf("Hello %s\n", name)
}

```

```go
//main.go
package main

import (
	"fmt"

	"example.com/myproject/pkg/calculate"
	"example.com/myproject/pkg/greeting"
	"example.com/myproject/pkg/helloworld"
)

func init() {
	fmt.Println("First init in main")
}

func main() {
	fmt.Println("Main function")

	// Call the SayHello function from the helloworld package
	helloworld.SayHello()
	greeting.SayHello()
	response := calculate.Add(1, 2)
	fmt.Printf("Calculate %d + %d = %d\n", 1, 2, response)
}
```

Bu durumda, dosyaların yüklenme sırası `main.go`, `calculate.go`, `greeting.go` ve `helloworld.go` olduğunu varsayarsak, `init` fonksiyonları da bu sırayla çağrılır. Bu sayede, her dosyanın `init` fonksiyonu belirli bir sırayla ve otomatik olarak çalıştırılır, böylece gerekli ayarlamalar ve kontroller belirli bir sıra ile gerçekleştirilmiş olur ve aşağıdaki çıktıyı elde ederiz.

```bash
First init in calculate package
First init in greeting package
First init in helloworld package
First init in main
Main function
Hello Veysel
Hello from greeting package
Calculate 1 + 2 = 3
```

Sonuç olarak, `init` fonksiyonu Go dilinde önemli bir fonksiyondur ve genellikle bir paketin veya dosyanın ilk başlatılması için kullanılır. Bu fonksiyon, paketin içindeki diğer tüm fonksiyonlardan önce otomatik olarak çağrılır ve herhangi bir giriş parametresi veya dönüş değeri olmaz. `init` fonksiyonu, paketin içinde birden fazla dosya bulunuyorsa ve her bir dosyanın bir `init` fonksiyonu varsa, dosyaların yüklenme sırasına göre çağrılır.# init-method-medium
