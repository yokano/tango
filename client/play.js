$(function() {
	// スクロール禁止
	$(window).bind('touchstart', function(event) {
		event.preventDefault();
	});
});