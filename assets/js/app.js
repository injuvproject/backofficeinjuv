$(document).ready(function() {
	"use strict";
	//carga la fecha cuando se crea una actividad
	$('#datetimepicker1').datetimepicker({
		format:'Y-m-d H:i',
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
    	console.log('date:' + $('#datetimepicker1').val())
		var formData = {'name': $('#name-activity').val(), 'description': $('#description-activity').val(), 'fechaExpiracion': $('#datetimepicker1').val(), 'recurso': $('#recursos').val(), 'estado': $('#estado').val(), 'pioridad': $('#pioridad').val(), 'adjunto': '0' }
			$.post( "http://localhost:3000/panel/nueva/actividad",
			 	formData, 
			 	function() {
			 		console.log('exito');
			})
  			.done(function(data, textStatus){
  					console.log('OK: true');
  					$('#myModal').modal('hide');
  					$('#name-activity').val("");
  					$('#description-activity').val("");
  					$('#datetimepicker1').val("");
  					alert('Se a cargado la tarea exitosamente');
  			})
	  		.fail(function(jqXHR, textStatus, errorThrown){
	    			console.log('error:'+errorThrown);
    				console.log('status:'+textStatus);
    				console.log('jqXHR:'+jqXHR);
    				
    				if (errorThrown === 'Bad Request'){
    					alert('Failed to connect to server');
    				}else if ( errorThrown === 'Unauthorized'){
    					alert('error guardar avtividad');

    				}
	  		})
	  		.always(function(data, textStatus, jqXHR){
	    		console.log( "finished" );
			});
	});


	
});