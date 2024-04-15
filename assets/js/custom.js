$(window).load(function () {

	"use strict";


	/* ========================================================== */
	/*   Portfolio                                            */
	/* ========================================================== */


	$(".show-project").on('click', function () {
		let targetId = $(this).data('target');
		$("#" + targetId).addClass('project-panel-show');
	})


	$(".close-project").on('click', function () {
		let targetId = $(this).data('target');
		$("#"+ targetId).removeClass('project-panel-show')
	})

	/* ========================================================== */
	/*   Hide Responsive Navigation On-Click                      */
	/* ========================================================== */

	$(".navbar-nav li a").on('click', function (event) {
		$(".navbar-collapse").collapse('hide');
	});


	/* ========================================================== */
	/*   Navigation Color                                         */
	/* ========================================================== */

	$('#navbar-collapse-02').onePageNav({
		filter: ':not(.external)'
	});


	/* ========================================================== */
	/*   SmoothScroll                                             */
	/* ========================================================== */

	$(".nav li a, a.scrool").click(function (e) {
		var full_url = this.href;
		var parts = full_url.split("#");
		var trgt = parts[1];
		var target_offset = $("#" + trgt).offset();
		var target_top = target_offset.top;

		$('html,body').animate({ scrollTop: target_top - 76 }, 1000);
		return false;

	});


	/* ========================================================== */
	/*   Newsletter                                               */
	/* ========================================================== */

	$('.newsletter_box .newsletter_form').each(function () {
		var form = $(this);
		//form.validate();
		form.submit(function (e) {
			if (!e.isDefaultPrevented()) {
				jQuery.post(this.action, {
					'email': $('input[name="nf_email"]').val(),
				}, function (data) {
					form.fadeOut('fast', function () {
						$(this).siblings('p.newsletter_success_box').show();
					});
				});
				e.preventDefault();
			}
		});
	});


	/* ========================================================== */
	/*   Register                                                 */
	/* ========================================================== */

	$('#register-form').each(function () {
		var form = $(this);
		//form.validate();
		form.submit(function (e) {
			if (!e.isDefaultPrevented()) {
				jQuery.post(this.action, {
					'names': $('input[name="register_names"]').val(),
					'phone': $('input[name="register_phone"]').val(),
					'email': $('input[name="register_email"]').val(),
					'ticket': $('select[name="register_ticket"]').val(),
				}, function (data) {
					form.fadeOut('fast', function () {
						$(this).siblings('p.register_success_box').show();
					});
				});
				e.preventDefault();
			}
		});
	})


	/* ========================================================== */
	/*   Contact                                                  */
	/* ========================================================== */
	$('#contact-form').each(function () {
		var form = $(this);
		//form.validate();
		form.submit(function (e) {
			if (!e.isDefaultPrevented()) {
				jQuery.post(this.action, {
					'names': $('input[name="contact_names"]').val(),
					'subject': $('input[name="contact_subject"]').val(),
					'email': $('input[name="contact_email"]').val(),
					'phone': $('input[name="contact_phone"]').val(),
					'message': $('textarea[name="contact_message"]').val(),
				}, function (data) {
					form.fadeOut('fast', function () {
						$(this).siblings('p').show();
					});
				});
				e.preventDefault();
			}
		});
	})

});

/* ========================================================== */
/*   Page Loader                                              */
/* ========================================================== */
$('#loader').fadeOut(100);