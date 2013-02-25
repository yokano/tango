$(function() {
	var word = $('#word');
	var addButton = $('#add');
	
	// 追加ボタンを押したら単語を追加
	addButton.bind('click', function() {
		$.ajax({
			url: '/add',
			data: {
				word: word.val()
			},
			success: function() {
				console.log('完成');
			},
			error: function() {
				console.log('ERROR');
			}
		});
	});
	
	// 単語一覧の表示
	
});