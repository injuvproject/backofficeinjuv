$(document).ready(function() {
	"use strict";
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