$(document).ready(function() {
	"use strict";
	//carga la fecha cuando se crea una actividad
	$('#datetimepicker1').datetimepicker({
		format:'d.m.Y H:i',
  		lang:'es'
	});
	//verifica si el checkbox que indica si es administrador es valido	
   	$("#admin").on( 'change', function() {
	    if( $(this).is(':checked') ) {
	        $(this).val("true") 
	    } else {
	         $(this).val("false") 
	    }
	});
   
   	$('.sortable').nestedSortable({
            handle: 'div',
            items: 'li',
            toleranceElement: '> div'
    });

    $('#btn-new-activity').click(function(){
		
		
		var formData = {'name': $('#name-activity').val().trim(), 'description': $('#description-activity').val().trim(), 'fechaExpiracion': $('#expiracion-activity').val().trim(), 'recurso': $('#recursos').val().trim(), 'estado': $('#estado').val().trim() };
			$.post( "htt",
			 	formData, 
			 	function() {
			 		console.log('exito')
			})
  			.done(function(data, textStatus) {
  					console.log('OK: true';
  					//
  			})
	  		.fail(function(jqXHR, textStatus, errorThrown) {
	    			console.log('error:'+errorThrown);
    				console.log('status:'+textStatus);
    				console.log('jqXHR:'+jqXHR);
    				
    				if (errorThrown === 'Bad Request'){
    					alert('Failed to connect to server')
    				}else if ( errorThrown === 'Unauthorized'){
    					alert('error guardar avtividad')

    				}
	  		})
	  		.always(function(data, textStatus, jqXHR) {
	    		console.log( "finished" );
			});
	
	});


	
});