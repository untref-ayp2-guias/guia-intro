# Guía: Introducción a Go

## Funciones

1. Implementar una función que, dado los coeficientes de un polinomio de grado
   $N$ (números flotantes), devuelva una representación en cadena de caracteres
   del polinomio completo.

   Por ejemplo, si recibe los coeficientes `3.0`, `-2.0` y `1.0` debe mostrar:

   ```text
   3.0-2.0x+1.0x^2
   ```

2. Implementar una función que reciba un número entero decimal y lo convierta a
   su representación en binario sin utilizar funciones propias del lenguaje que
   lo hagan directamente.

3. Escribir una función que reciba una lista de enteros y devuelva el menor y el
   mayor de la lista.

## Bucles

1. Escribir una función que recibe un número entero `n` (no negativo), y
   devuelve su factorial.

2. Escribir una función que recibe dos enteros: `a` y `b` y devuelve el
   producto. Para el cálculo, utilizar sumas sucesivas.

3. Escribir una función que indique si un número entero, ingresado por un
   usuario, es primo (devuelve verdadero o falso).

## Arreglos

1. Escribir una función que reciba un arreglo de enteros como parámetros y
   devuelva la sumatoria de todos sus elementos.

2. Escribir una función que reciba dos arreglos (que representan vectores) del
   mismo tamaño y devuelvan la suma vectorial y el producto escalar de los
   vectores.

3. Escribir una función que reciba dos arreglos _A_ y _B_, de _N_ y _M_
   elementos respectivamente que representan conjuntos de enteros y devuelva una
   arreglo con la unión y otro con la intersección de _A_ y _B_.

## Punteros

1. Escribir una función con la siguiente firma `func Swap(px, py *int)` que
   permita intercambiar el valor almacenado en dos variables enteras a través de
   punteros.

## Estructuras

1. Implementar la estructura `Circulo` similiar a las vistas en el módulo
   `figuras` del taller de Go. La estructura debe tener un campo `radio` de tipo
   entero y un campo `centro` de tipo `Punto`.

   `Circulo` debe implementar los métodos de la interface `Figura`.

---

## Como trabajar con la guía

Una vez clonado el repositorio primero debemos asegurarnos de que las
dependencias de la guía están instaladas.

```bash
go mod tidy
```

Para asegurarnos de que nuestro código no tiene errores de compilación
ejecutamos:

```bash
go build ./...
```

Para ejecutar todos los tests de la guía ejecutamos:

```bash
go test ./...
```

Si queremos ejecutar un test en particular, podemos hacerlo de la siguiente
manera:

```bash
go test ./... -run ^TestSumatoria$
```

donde `TestSumatoria` es el nombre del test.
