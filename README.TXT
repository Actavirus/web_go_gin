https://fokusov.com/posts/razrabotka-web-prilozhenij-i-mikroservisov-na-golang-s-gin/
Мы рас­смот­рим со­зда­ние про­ек­та и сбор­ку неслож­но­го при­ло­же­ния с Gin, ко­то­рое бу­дет вы­во­дить спи­сок то­пи­ков и от­дель­ный то­пик.
На­ше при­ло­же­ние — это про­стой ме­не­джер то­пи­ков. Оно долж­но:
поз­во­лять поль­зо­ва­те­лям ре­ги­стри­ро­вать­ся с ло­ги­ном и па­ро­лем (для неав­то­ри­зо­ван­ных поль­зо­ва­те­лей),
поз­во­лять поль­зо­ва­те­лям ав­то­ри­зо­вать­ся (для неав­то­ри­зо­ван­ных поль­зо­ва­те­лей),
поз­во­лять поль­зо­ва­те­лям за­вер­шать се­анс (для ав­то­ри­зо­ван­ных поль­зо­ва­те­лей),
поз­во­лять поль­зо­ва­те­лям со­зда­вать то­пи­ки (для ав­то­ри­зо­ван­ных поль­зо­ва­те­лей),
Вы­во­дить спи­сок всех то­пи­ков на глав­ной стра­ни­це (для всех поль­зо­ва­те­лей), и
Вы­во­дить то­пик на его соб­ствен­ной стра­ни­це (для всех поль­зо­ва­те­лей). Вдо­ба­вок к это­му мы сде­ла­ем, что­бы спи­сок то­пи­ков и от­дель­ные то­пи­ки бы­ли до­ступ­ны в фор­ма­тах HTML, JSON и XML.