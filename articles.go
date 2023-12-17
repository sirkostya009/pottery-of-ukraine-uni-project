package main

type Image struct {
	Title string
	Url   string
}

var descriptions = map[string]string{
	"Горщик":   "Горщик, як символ українського гончарства, втілює багатовікову традицію ремесла. Його форма та декор відображають національні особливості та красу української культури. Цей предмет виготовляється вручну з використанням природних матеріалів, що підкреслює екологічність та автентичність українського гончарства. Горщик – це не лише побутовий предмет, а й шедевр ремесла, який зберігає спадщину та віддзеркалює мистецтво українських майстрів.",
	"Глек":     "Глек, утворений в руках українських гончарів, представляє собою елегантний синтез мистецтва та корисності. Його форма та орнаментика народжуються з глини, пройнятої духом української традиції. Глек відзначається не лише вишуканістю дизайну, але й практичністю, роблячи його ідеальним для зберігання та подачі напоїв. Ручна робота та використання натуральних матеріалів підкреслюють автентичність та високу якість українського гончарства. Глек стає не просто посудом, а витвором мистецтва, який зберігає дух національної культури.",
	"Макітра":  "Макітра, виготовлена українськими гончарями, є невід'ємною частиною культурної спадщини. Її унікальна форма та прикраси вражають своєрідністю та красою. Майстри вкладають у свою роботу не лише вміння, а й дух споконвічних традицій. Ручна робота та використання найкращих матеріалів надають макітрі особливу вартість та створюють унікальний характер. Макітра стає не лише предметом одягу, але і витвором мистецтва, який підкреслює велич і красу української національної традиції.",
	"Куманець": "Куманець, вироблений в руках українських гончарів, це не просто предмет, а витвір мистецтва, який вражає своєю унікальністю та експресією. Його форма і декор створюють гармонійний образ, де відчувається дух української культури та традицій. Ручна робота гончара, використання високоякісних матеріалів та врочистість вигляду роблять куманець важливим символом національної спадщини. Кожен елемент має свою унікальну історію, а спільний образ куманця є втіленням величі та глибини української ідентичності.",
	"Чайник":   "Чайник, створений у руках вправного українського гончара, – це справжнє шедевр гончарного мистецтва, що вражає своєю елегантністю та функціональністю. Його форма та дизайн тонко відтіняють витончений смак та вражають неповторною красою. Виготовлений вручну з використанням високоякісних матеріалів, чайник не лише виконує практичну роль у приготуванні напоїв, але і стає предметом, який виражає національну гордість та майстерність. Ретельна увага до деталей та гармонійне поєднання форми і функції роблять кожен чайник унікальним твором, який надихає та перетворює звичайний акт чаювання у витончений ритуал.",
}

var Glossary = []Image{{
	Title: "Горщик",
	Url:   "/static/горщик.svg",
}, {
	Title: "Глек",
	Url:   "/static/глек.svg",
}, {
	Title: "Свічки",
	Url:   "/static/свічки.svg",
}, {
	Title: "Свічка",
	Url:   "/static/свічка.svg",
}, {
	Title: "Баран",
	Url:   "/static/баран.svg",
}, {
	Title: "Декор",
	Url:   "/static/декор.svg",
}, {
	Title: "Макітра",
	Url:   "/static/макітра.svg",
}, {
	Title: "Чайник",
	Url:   "/static/чайник.svg",
}, {
	Title: "Куманець",
	Url:   "/static/куманець.svg",
}, {
	Title: "Квочка",
	Url:   "/static/квочка.svg",
}}
