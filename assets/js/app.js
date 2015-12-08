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

	
});