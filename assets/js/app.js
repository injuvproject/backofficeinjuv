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
   
    $('#btn-new-activity').click(function(){
    	
		var formData = {'name': $('#name-activity').val(), 'description': $('#description-activity').val(), 'fechaExpiracion': $('#datetimepicker1').val(), 'recurso': $('#recursos').val(), 'estado': $('#estado').val(), 'pioridad': $('#pioridad').val(), 'adjunto': '0' }
			$.post( "http://188.166.41.225:3000/panel/nueva/actividad",
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
  					location.reload();
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

	$('#btn-update-activityIM').click(function(){
		console.log("id:" + $('#id-activityIM').val());
		console.log("Estado:" + $('#estadoIM').val());
		var formData = {'id': $('#id-activityIM').val(), 'estado': $('#estadoIM').val()}
			$.post( "http://188.166.41.225:3000/panel/actualiza/actividad",
			 	formData, 
			 	function() {
			 		console.log('exito');

	

			})
  			.done(function(data, textStatus){
  					console.log('OK: true');
  					$('#activityModalIM').modal('hide');
  					$('#id-activityIM').val("");
  					$('#estadoIM').val("");
  					
  					alert('Se a actuaizado la tarea exitosamente');
  					location.reload();


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

	$('#btn-update-activityT').click(function(){
		console.log("id:" + $('#id-activityT').val());
		console.log("Estado:" + $('#estadoT').val());
		var formData = {'id': $('#id-activityT').val(), 'estado': $('#estadoT').val()}
			$.post( "http://188.166.41.225:3000/panel/actualiza/actividad",
			 	
			 	formData, 
			 	function() {
			 		console.log('exito');
			})
  			.done(function(data, textStatus){
  					console.log('OK: true');
  					$('#activityModalTerminados').modal('hide');
  					$('#id-activityT').val("");
  					$('#estadoT').val("");
  					
  					alert('Se a actuaizado la tarea exitosamente');
  					location.reload();
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

	
	$('#btn-update-activityP').click(function(){
		console.log("id:" + $('#id-activityP').val());
		console.log("Estado:" + $('#estadoP').val());
		var formData = {'id': $('#id-activityP').val(), 'estado': $('#estadoP').val()}
			$.post( "http://188.166.41.225:3000/panel/actualiza/actividad",
			 	formData, 
			 	function() {
			 		console.log('exito');
			})
  			.done(function(data, textStatus){
  					console.log('OK: true');
  					$('#activityModalProceso').modal('hide');
  					$('#id-activityP').val("");
  					$('#estadoP').val("");
  					
  					alert('Se a actuaizado la tarea exitosamente');
  					location.reload();
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


	$('#btn-update-activityPEN').click(function(){
		console.log("id:" + $('#id-activityPEN').val());
		console.log("Estado:" + $('#estadoPEN').val());
		var formData = {'id': $('#id-activityPEN').val(), 'estado': $('#estadoPEN').val()}
			$.post( "http://188.166.41.225:3000/panel/actualiza/actividad",
			 	formData, 
			 	function() {
			 		console.log('exito');
			})
  			.done(function(data, textStatus){
  					console.log('OK: true');
  					$('#activityModalPendientes').modal('hide');
  					$('#id-activityPEN').val("");
  					$('#estadoPEN').val("");
  					
  					alert('Se a actuaizado la tarea exitosamente');
  					location.reload();
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
	
	
	
	
	
	 $('#btn-closeImpedidos').click(function(){
	 	$('#myModalLabel').val("");
  		$('#description-impedidos').val("");
  		$('#date-impedido').val("");
  		$('#myModalLabel').attr('h4', '');

	 });


	
});