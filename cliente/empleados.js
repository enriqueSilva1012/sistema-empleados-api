const API="http://localhost:8080/empleados";

async function cargarEmpleados(){

const respuesta=
await fetch(API);

const empleados=
await respuesta.json();

const tabla=
document.getElementById(
"tablaEmpleados"
);

tabla.innerHTML="";

empleados.forEach(
empleado=>{

tabla.innerHTML += `

<tr>

<td>${empleado.id}</td>

<td>${empleado.nombre}</td>

<td>${empleado.apellido}</td>

<td>${empleado.puesto}</td>

<td>${empleado.salario}</td>

<td>${empleado.departamento}</td>

<td>${empleado.estado}</td>

<td>

<button>
Editar
</button>

<button>
Eliminar
</button>

</td>

</tr>

`;

});

}

async function guardarEmpleado(){

const empleado={

nombre:
document.getElementById(
"nombre"
).value,

apellido:
document.getElementById(
"apellido"
).value,

puesto:
document.getElementById(
"puesto"
).value,

salario:
parseFloat(
document.getElementById(
"salario"
).value
),

departamento:
document.getElementById(
"departamento"
).value,

fechaIngreso:
document.getElementById(
"fechaIngreso"
).value,

estado:
document.getElementById(
"estado"
).value

};

await fetch(
API,
{

method:"POST",

headers:{

"Content-Type":
"application/json"

},

body:
JSON.stringify(
empleado
)

}

);

cargarEmpleados();

}

document
.getElementById(
"guardar"
)
.addEventListener(
"click",
guardarEmpleado
);

cargarEmpleados();