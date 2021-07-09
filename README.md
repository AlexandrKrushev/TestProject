<h1><a id="__0"></a>Тестовое задание</h1>
<p>Реализовать throttler-обёртку для типа Transport из стандартной библиотеки (<a href="https://golang.org/pkg/net/http/#Transport">https://golang.org/pkg/net/http/#Transport</a>).<br>
Обёртка должна реализовывать интерфейс RoundTripper (<a href="https://golang.org/pkg/net/http/#RoundTripper">https://golang.org/pkg/net/http/#RoundTripper</a>) и инициализироваться следующими параметрами:</p>
<ul>
<li>RoundTripper, который будет оборачиваться</li>
<li>Лимит запросов в единицу времени (целое число, если равно 0 то throttling не применяется)</li>
<li>Единица времени учёта (тип time.Duration (<a href="https://golang.org/pkg/time/#Duration">https://golang.org/pkg/time/#Duration</a>))</li>
<li>Список префиксов исключений URL для которых throttling не будет задействован (если список пуст или nil - то исключений нет)</li>
<li>Флаг быстрого возврата ошибки</li>
</ul>
<h3><a id="_10"></a>Примечания</h3>
<p>Если частота запросов превышает лимит, то запрос должен быть отложен до момента когда его выполнение не вызовет превышение лимита либо завершён со специальной ошибкой (в зависимости от флага быстрого возврата ошибки).<br>
Для учёта частоты запросов можно считать что они выполняются мгновенно, коды возврата не имеют значения, запросы не подпадающие под условия фильтров не учитываются.<br>
Списки префиксов URL могут содержать * в любой части пути.<br>
Нужно помнить что обёртка может использоваться из многих параллельных горутин, а так же может быть использована в цепочке из нескольких обёрток.<br>
Решение задачи должно быть оформлено в виде отдельного репозитория на гитхабе.</p>
<h3><a id="__17"></a>Пример использования:</h3>
<pre><code>throttled := NewThrottler(
    http.DefaultTransport,
    60,
    time.Minute, // 60 rpm
    []string{&quot;/servers/*/status&quot;, &quot;/network/&quot;}, // except servers status and network operations
    false, // wait on limit
)

client := http.Client{
    Transport: throttled,
}

// ...
// no throttling
resp, err:= client.GET(&quot;http://apidomain.com/network/routes&quot;) 
// ...

// ... 
// throttling might be used
req := http.NewRequest(&quot;PUT&quot;, &quot;http://apidomain.com/images/reload&quot;, nil)
resp, err:= client.Do(req) 
// ...

// ...
// no throttling
resp, err:= client.GET(&quot;http://apidomain.com/servers/1337/status?simple=true&quot;) 
// ...
</code></pre>

</body></html>
