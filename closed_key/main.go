/* Во всех крупных IT-компаниях немалое внимание уделяется вопросам информационной безопасности, и Яндекс не является исключением.
Дима и Егор разрабатывают новый сервис YD (Yandex Dorogi) и в данный момент занимаются аудитом его безопасности. Для шифрования пользовательских данных в YD используется алгоритм шифрования с открытым ключом YS (Yandex Shifrovatel).
Схема работы алгоритма YS такова: для каждого сервиса генерируется закрытый ключ (p,q), где p и q — натуральные числа. По закрытому ключу (p,q) генерируется открытый ключ (НОД(p,q), НОК(p,q)), 
который доступен всем пользователям. Если злоумышленник сможет по открытому ключу получить закрытый ключ, то он получит доступ ко всем данным YD и нанесёт сервису непоправимый вред.
Конечно же, Егор и Дима не хотят этого допустить, поэтому они хотят сделать так, чтобы злоумышленнику пришлось перебрать очень много вариантов открытого ключа, прежде чем он сможет его угадать.

Дима уже сгенерировал закрытый ключ для YD и получил на его основе открытый ключ (x,y). 
Егору сразу же стало интересно, сколько вариантов закрытого ключа придётся перебрать злоумышленнику для взлома YD в худшем случае, иными словами, сколько существует закрытых ключей (p,q) таких, что открытым ключом для них является (x,y).
К сожалению, у Егора есть много других задач, очень важных для запуска YD, поэтому он просит вас вычислить это количество за него.
 */
package main

import (
	"fmt"
)

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func CountPairs(g, l int) int {
	if l%g != 0 {
		return 0
	}

	k := l / g
	count := 0

	for m := 1; m*m <= k; m++ {
		if k%m == 0 {
			n := k / m
			if Gcd(m, n) == 1 {
				if m == n {
					count++
				} else {
					count += 2
				}
			}
		}
	}
	return count
}
func main() {
	var g, l int
	fmt.Scan(&g, &l)
	result := CountPairs(g, l)
	fmt.Print(result)
}
