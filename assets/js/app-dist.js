$(document).ready(function(){"use strict";$("#datetimepicker1").datetimepicker({format:"d.m.Y H:i",lang:"es"}),$("#admin").on("change",function(){$(this).val($(this).is(":checked")?"true":"false")}),$(".sortable").nestedSortable({handle:"div",items:"li",toleranceElement:"> div"}),$("#btn-new-activity").click(function(){var e={name:$("#name-activity").val(),description:$("#description-activity").val(),fechaExpiracion:$("#expiracion-activity").val(),recurso:$("#recursos").val(),estado:$("#estado").val(),pioridad:$("#pioridad").val(),adjunto:"0"};$.post("http://localhost:3000/panel/nueva/actividad",e,function(){console.log("exito")}).done(function(e,t){console.log("OK: true")}).fail(function(e,t,o){console.log("error:"+o),console.log("status:"+t),console.log("jqXHR:"+e),"Bad Request"===o?alert("Failed to connect to server"):"Unauthorized"===o&&alert("error guardar avtividad")}).always(function(e,t,o){console.log("finished")})})});