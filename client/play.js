/**
 * 単語学習画面のスクリプト
 * var words に単語配列がサーバによって挿入されている
 */
$(function() {
	// スクロール禁止
	$(window).bind('touchstart', function(event) {
		event.preventDefault();
	});
	
	
//	wordTemplate = `
//		<div data-role="page" id="%d">
//			<div data-role="content">
//				<div class="word">%s</div>
//				<div class="meaning">%s</div>
//			</div>
//			<div class="trashbox">
//				<img src="/client/trashbox.png" />
//			</div>
//		</div>
//	`
});