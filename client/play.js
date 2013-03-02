/**
 * 単語学習画面のスクリプト
 * var words に単語配列がサーバによって挿入されている
 */
$(function() {
	// スクロール禁止
	$(window).bind('touchstart', function(event) {
		event.preventDefault();
	});
	
	// Microsoft Translate API へ接続
	$.ajax({
		url: 'http://api.microsofttranslator.com/V2/Ajax.svc/GetAppIdToken',
		method: 'get',
		datatype: 'json',
		error: function() {
			console.log('error');
		},
		success: function(data) {
			console.log(data);
		}
	});
});