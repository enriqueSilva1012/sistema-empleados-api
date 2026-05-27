const API="http://192.168.1.255:8080/empleados";


async function cargarEmpleados(){

    try{

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
        (empleado)=>{

            tabla.innerHTML+=`

            <tr>

            <td>

            ${empleado.ID}

            </td>

            <td>

            ${empleado.nombre}

            </td>

            <td>

            ${empleado.puesto}

            </td>

            <td>

            ${empleado.departamento}

            </td>

            <td>

            ${
            empleado.fechaIngreso
            ?.split("T")[0]
            }

            </td>

            </tr>

            `;

        });

    }

    catch(error){

        console.log(error);

    }

}



async function guardarEmpleado(){


    const empleado={

        nombre:
        document.getElementById(
            "nombre"
        ).value,

        puesto:
        document.getElementById(
            "puesto"
        ).value,

        departamento:
        document.getElementById(
            "departamento"
        ).value,

        fechaIngreso:
        document.getElementById(
            "fechaIngreso"
        ).value
        +"T00:00:00Z"

    };


    try{

        const respuesta=
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

        const datos=
        await respuesta.json();


        if(
        !respuesta.ok
        ){

            alert(
            JSON.stringify(
            datos
            )
            );

            return;

        }


        alert(
        "Registro exitoso"
        );

        cargarEmpleados();

    }

    catch(error){

        console.log(
        error
        );

    }

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