/* Вы обслуживаете сайт и отслеживаете возникающие проблемы. Клиент получил ошибку после того, как попытался добавить свой пост в систему. Вы хотите понять, на каком из серверов эта ошибка произошла.
Есть n серверов, на i-й из них приходится a процентов запросов, из которых b процентов завершаются с ошибкой. Для каждого сервера найдите вероятность того, что ошибка произошла именно на нём. */


#include <iostream>
#include <iomanip>

using namespace std;

int main() {
    int n;
    cin >> n;
    double prob[n][2];
    double total = 0.0;
    for (int i = 0; i < n; i++) {
        double a, b;
        cin >> a >> b;
        prob[i][0] = a;
        prob[i][1] = b;
        total += (a * b) / 100;
    }
    for (int i = 0; i < n; i++) {
        cout << fixed << setprecision(12) << prob[i][0] * prob[i][1] / 100 / total << endl;
    }
}