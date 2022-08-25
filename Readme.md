
<a name="readme-top"></a>




<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
-->
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="https://pbs.twimg.com/profile_images/1524162331294543874/ikWkgx1g_400x400.jpg" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">PRUEBA DIEGO CHAMORRO</h3>

  <p align="center">
    Mi reto de aprendizaje.
    <br />Aspirante al cargo: Senior Software Engineer - Tech Lead.
    <br />
    <a href="https://github.com/dchamorros/Testruct"><strong>Explorar proyecto »</strong></a>
    <!-- <br />
    <br />
    <a href="https://github.com/github_username/repo_name">View Demo</a>
    ·
    <a href="https://github.com/github_username/repo_name/issues">Report Bug</a>
    ·
    <a href="https://github.com/github_username/repo_name/issues">Request Feature</a> -->
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Tabla de Contenido</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## Acerca del Proyecto

Esta es una implementacion de un estructura T, de forma N Dimensional con valores float 32, sobre la cual se tiene implementado las operaciones: `Reshape`, `Hadamard Product`, `Index Select`


<p align="right">(<a href="#readme-top">Ir arriba</a>)</p>



### Desarrollado con

* GO

### Nota

* Este proyecto corresponde a mi primera implementación en Go. Y puedo decir que me pareció facil de aprender.

<p align="right">(<a href="#readme-top">Ir arriba</a>)</p>



<!-- GETTING STARTED -->
## Como Empezar

Para ejecutar este proyecto, primero ten una copia local y corre los ejemplos, siguiendos los pasos a continuación:


### Instalación

1. Conectarse a [https://github.com/dchamorros/Testruct.git](https://github.com/dchamorros/Testruct.git)
2. Clonar the repositorio
   ```sh
   git clone https://github.com/dchamorros/Testruct.git
   ```

<p align="right">(<a href="#readme-top">Ir arriba</a>)</p>



<!-- USAGE EXAMPLES -->
## Uso

La implementación de las funciones incluye métodos para hacer las respectivas pruebas. Sin embargo, antes de usar los ejemplos, es necesario entender cómo funciona la estructura T.

### shape
Es un array de enteros que representa las dimensiones de la estructura T.
 
 `Ejemplo 1 Dimensión:`
   ```sh
   shape := []int64{2}
   ```
   En este caso, el shape indica que la estructura T va a ser de 1 dimensión dado que el arreglo tiene logitud 1, e indica que en esa dimensión habrá un array con 2 elementos.

   `Ejemplo N Dimensiones:`
  ```sh
   shape := []int64{3, 2, 1, 4}
   ```
   En este caso, el shape indica que la estructura T va a ser de 4 dimensiones dado que el arreglo tiene logitud 4, e indica que:
   - En la cuarta dimensión habrá un array con 4 elementos de tipo float. 
   - En la dimesión 3 habrá un array con 1 elemento.
   - En la dimension 2 habrá un array con 2 elementos.
   - Y, en la dimensión 1 habrá un array con 3 elementos.

### datos
Hay 2 formas de suministrarle datos a la estructura T:

- Generación aleatoria de valores
- Fijándo valores

 
 `Ejemplo  Generación aleatoria de Valores:`
   ```sh
   var myStructT t.StructT
	myStructT.Shape = shape
	myStructT.SetData(nil)
   ```
   Si desea que el sistema genere valores aleatorios para la estructura definida anteriorme se le deberá pasar un valor `nil` en la inicialización de la data.


   `Ejemplo Fijándo valores:`
   ```sh
    data := [][]int{{1, 2}, {3, 4}}

    var myStructT t.StructT
	myStructT.Shape = shape
	myStructT.SetData(data)
   ```
   Si desea que fijar sus propios valores a la estructura definida anteriorme se le deberá pasar un valor `data` en la inicialización de la data de cada uno de los ejemplos.


### operaciones
- hadamard

   ```sh
    shape := []int64{2, 2}
	dataA := [][]int{{1, 2}, {3, 4}}
	dataB := [][]int{{2, 2}, {4, 4}}
	
    var myStructA t.StructT
	myStructA.Shape = shape
	myStructA.SetData(dataA)
	
	var myStructB t.StructT
	myStructB.Shape = shape
	myStructB.SetData(dataB)
	
	fmt.Println("Estructura Resultado:", myStructA.HadamardProduct(myStructA, myStructB))
   ```
    `ejemplos disponibles:`
    1. sampleHadamard2X2()
    2. sampleHadamardND()

- reshape

   ```sh
   newShape := []int64{2, 1}
	fmt.Println("Aplicando ReShape:", newShape)
	myStructT.Reshape(newShape)
   ```
   Dentro de las funciones de ejemplo que se encuentran en el código, deberá definirse un newShape con las dimensiones de la estructura resultante de la operación, considerando que debe mantener la misma cantidad de elementos para que la operación pueda darse.

   `ejemplos disponible:`
    1. sampleReShape1D()
    2. sampleReShape2D()
    3. sampleReShapeND()


- indexSelec

   ```sh
    dim := 1
	indexVector := []int64{0}
	var myStructT t.StructT
	myStructT.Shape = shape
	myStructT.SetData(data)
	fmt.Println("Estructura Resultado:", myStructT.IndexSelect(dim, indexVector))
   ```
En cada una de las funciones de ejemplo de la operacion indexSelect que desee probar, se debe:
1. Definir la dimensión desde la cual se tomarán los datos en el atributo `dim`.
2. Definir el vector de indices en el atributo `indexVector`, especiaficandose los indices que se desean considerar en la operación.
3. Fijarle a la estructura `myStructT` la data previamente definida.
4. Llamar la función indexSelect de la estructura `myStructT` pasandole los parametros: `dim` e `indexVector`.

`Ejemplos disponibles:`

Para hacer uso de las funciones `indexSelect` abra el archivo hello.go ubicado en la carpeta src. Aquí encontrará la siguiente lista de ejemplos:

1. sampleIndexSelect1()
2. sampleIndexSelect4()
3. sampleIndexSelectN()




<!-- CONTACT -->
## Contacto

Diego Chamorro Salas - [@dchamorros](https://twitter.com/dchamorros) - dchamorros@gmail.com

<p align="right">(<a href="#readme-top">Ir arriba</a>)</p>





<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/github_username/repo_name.svg?style=for-the-badge
[contributors-url]: https://github.com/github_username/repo_name/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/github_username/repo_name.svg?style=for-the-badge
[forks-url]: https://github.com/github_username/repo_name/network/members
[stars-shield]: https://img.shields.io/github/stars/github_username/repo_name.svg?style=for-the-badge
[stars-url]: https://github.com/github_username/repo_name/stargazers
[issues-shield]: https://img.shields.io/github/issues/github_username/repo_name.svg?style=for-the-badge
[issues-url]: https://github.com/github_username/repo_name/issues
[license-shield]: https://img.shields.io/github/license/github_username/repo_name.svg?style=for-the-badge
[license-url]: https://github.com/github_username/repo_name/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/dchamorros
[product-screenshot]: images/screenshot.png
[Next.js]:https://go.dev/images/go-logo-white.svg-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 
