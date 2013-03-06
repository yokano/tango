$(function() {
	var word = $('#word');
	var addButton = $('#add');
	var clearButton = $('#clear');
	var words = $('ul');
	
	// 追加ボタンを押したら単語を追加
	addButton.bind('click', function() {
	
		// 入力チェック
		if(word.val() == '') {
			return;
		}
		
		// サーバへ送信
		$.ajax({
			url: '/add',
			data: {
				word: word.val()
			},
			dataType: 'json',
			async: false,
			success: function(data) {
				if(data.wordnum > 1) {
					words.html(words.html() + ', ' + word.val());
				} else {
					words.html(word.val());
				}
				word.val('');
			},
			error: function() {
				console.log('ERROR');
			}
		});
	});
	
	// 全消しボタン
	clearButton.bind('click', function() {
		if(window.confirm('単語をすべて削除しますか？')) {
			$.ajax({
				url: '/clear',
				async: false,
				success: function() {
					words.empty();
				},
				error: function() {
					console.log('ERROR');
				}
			});
		}
	});
});