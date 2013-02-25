$(function() {
	var word = $('#word');
	var addButton = $('#add');
	
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
			async: false,
			success: function() {
				$('ul').html($('ul').html() + ', ' + word.val());
				word.val('');
			},
			error: function() {
				console.log('ERROR');
			}
		});
	});
});