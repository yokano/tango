/**
 * 単語学習画面のスクリプト
 * var words に単語配列がサーバによって挿入されている
 * @property {Number} current 現在のページ
 */
$('#play').bind('pagebeforeshow', function() {
	var current = 0;
	var pages = [];
	
	// スクロール禁止
	$(this).bind('touchstart', function(event) {
		event.preventDefault();
	});
	
	// CSSロード
	$('<link rel="stylesheet" href="/client/play.css"></link>').appendTo($('body'));
	
	// 単語ページ作成
	for(var i = 0; i < words.length; i++) {
		pages[i] = $('\
			<div data-role="page" id="page' + i + '" class="wordpage">\
				<div data-role="content">\
					<div class="word">' + words[i].Word + '</div>\
					<div class="meaning">' + words[i].Meaning + '</div>\
					<div class="trashbox">\
						<img src="/client/trashbox.png" />\
					</div>\
				</div>\
			</div>\
		').appendTo($('body'));
		pages[i].open = false;
	}
	
	// 終了ページ作成
	pages[i] = $('\
		<div data-role="page" id="page' + i + '">\
			<div>すべての単語が終わりました</div>\
			<div data-role="button" data-icon="refresh" id="replay">もう１度</div>\
			<a href="/" data-role="button" data-icon="home" id="exit">やめる</a>\
		</div>\
	').appendTo($('body'));
	$('#replay').bind('tap', function() {
		for(var i = 0; i < pages.length - 1; i++) {
			close(i);
		}
		current = 0;
		$.mobile.changePage(pages[current]);
	});
	
	// スワイプでページ切り替え
	$('.wordpage').bind('swipeleft', function() {
		if(current <= words.length) {
			current++;
			$.mobile.changePage(pages[current], {
				transition: 'slide'
			});
		}
	});
	$('.wordpage').bind('swiperight', function() {
		if(current > 0) {
			current--;
			$.mobile.changePage(pages[current], {
				transition: 'slide',
				reverse: true
			});
		}
	})
	
	// 画面タップで意味表示
	var open = function(pageID) {
		if(!pages[pageID].open) {
			pages[pageID].open = true;
			pages[pageID].find('.meaning').fadeIn('slow');
			pages[pageID].find('.trashbox').fadeIn('slow');
			pages[pageID].children().animate({
				marginTop: '-50px'
			});
		}
	};
	var close = function(pageID) {
		if(pages[pageID].open) {
			pages[pageID].open = false;
			pages[pageID].find('.meaning').hide();
			pages[pageID].find('.trashbox').hide();
			pages[pageID].children().animate({
				marginTop: '0px'
			});
		}
	};
	$('.wordpage').bind('tap', function() {
		open(current);
	});
	
	// ゴミ箱をタップしたらページを削除
	$('.trashbox').bind('tap', function(e) {
		e.stopPropagation();
		e.preventDefault();
		pages.splice(current, 1);
		$.mobile.changePage(pages[current]);
	});
	
	$.mobile.changePage(pages[current]);
});