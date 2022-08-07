package tasks

/*
 Задача 15.	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
 Приведите корректный пример реализации.
 Код, который нужно исправить:
 var justString string

 func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
 }
 func main() {
	someFunc()
 }
*/
// Использовать глобальные переменные - плохая практика!
import "fmt"

func TaskFifteen() []byte {
	justString := `ko qwwyefhwehf8 832 jhu	u89U Q	89 ujH Dpqh hhjhdwh udhwudh   widjwndjshduwu 
iewj di jwfijwpdhw8cuwj2837r832u rjefjwneknwejd dsjcuhv rhe
оывмвосрыгурствс ввшсшвосоывосш лоыыовгмсцвосыоо ji ejwfljeofij8u j.`

	hugeStringByte := []byte(justString)
	justStringByte := hugeStringByte[:100]        // этот срез, для того, чтобы указать диапазон данных, например [23:28]
	newSlice := make([]byte, len(justStringByte)) // создаем пустой срез с заданной длиной
	copy(newSlice, hugeStringByte)                // копируем и возвращаемся из функции
	fmt.Println(newSlice)
	return newSlice
}
