$(function() {
	var word = $('#word');
	var addButton = $('#add');
	var clearButton = $('#clear');
	var words = $('#words');
	
	// 追加ボタンを押したら単語を追加
	addButton.bind('tap', function(e) {
				
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
			async: false,
			dataType: 'json',
			success: function(data) {
				$('<li>' + word.val() + '</li>').appendTo(words);
				words.listview('refresh');
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
				}
			});
		}
	});
});