Natas5

login:natas5
password:iX6IOfmpN7AYOQGPwtn3fXpbaJVJcHfq

1)На главной странице видим сообщение о том что мы не авторизованы, открываем burp и смотрим как выглядит запрос.
![изображение](https://user-images.githubusercontent.com/54982244/173063836-ae1056c0-1eec-49f1-bddf-c8abc1db2940.png)
2)Перехваченный запрос выгляит так.
![изображение](https://user-images.githubusercontent.com/54982244/173063970-28a2f2f3-2719-4471-8774-3e15d67e661e.png)
3)Обращаем внимаение на Cookie: loggedin=0, меняем на Cookie: loggedin=1 и отправляем запрос.
![изображение](https://user-images.githubusercontent.com/54982244/173064103-9922dc68-5d1f-492c-96a2-8f5e4d6e35e7.png)
4)Забираем пароль для следующего уровня
![изображение](https://user-images.githubusercontent.com/54982244/173064155-aa30df3a-27a2-45d7-a8b5-94ad8a1e0926.png)

2 вариант.
1)Cookie можно поменять не перехватывая сам запрос, открываем страницу, ПКМ + Иследовать. Отрываем панель разработчика. Переходим в Хранилище, Cookie.
Видим наш loggedin со значением 0.
![изображение](https://user-images.githubusercontent.com/54982244/173065445-831fa45e-9867-4cf6-9b46-48a5e786765c.png)
2)Меняем значение на 1 и презагружаем страницу. Забираем пароль для следующего уровня.
![изображение](https://user-images.githubusercontent.com/54982244/173065633-63e832ac-8525-4bfa-848c-120a1a93d06b.png)
