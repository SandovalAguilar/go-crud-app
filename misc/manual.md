# Manual de Comandos Esenciales para Ubuntu Server

-----

## Acceso y Conexión

### 1\. En la Consola del Servidor (Local)

Cuando accedes directamente a la máquina de Ubuntu Server, verás la siguiente solicitud en pantalla.

1.  **Introduce tu Nombre de Usuario:**
    ```text
    Ubuntu 22.04 LTS ip-xxx-xxx-xxx-xxx tty1

    ip-xxx-xxx-xxx-xxx login: **[Escribe tu nombre de usuario]**
    ```
2.  **Introduce tu Contraseña:**
    ```text
    Password: **[Escribe tu contraseña (no verás los caracteres)]**
    ```

Si las credenciales son correctas, se te presentará el *prompt* (símbolo del sistema).

-----

### 2\. Vía SSH (Acceso Remoto)

La forma estándar de acceder a un servidor es de forma remota, utilizando el protocolo **SSH (Secure Shell)**.

El comando general es:

```bash
ssh [nombre_de_usuario]@[dirección_IP_del_servidor]
```

#### Para Sistemas Unix-like (Linux/macOS)

Si usas Linux o macOS, el cliente SSH viene **instalado por defecto**. Abre la **Terminal** e ingresa el comando directamente:

```bash
ssh tu_usuario@192.168.1.50
```

#### Para Sistemas Windows

En Windows, tienes dos opciones principales:

1.  **PowerShell / Símbolo del sistema (CMD):** Las versiones modernas de Windows (10 y 11) incluyen el cliente SSH de forma nativa.
    ```bash
    ssh tu_usuario@192.168.1.50
    ```
2.  **Cliente PuTTY (GUI):** Para versiones anteriores o si prefieres una **interfaz gráfica**.
      * Descarga e instala **PuTTY**.
      * Abre PuTTY e introduce la **Dirección IP** del servidor en el campo "Host Name (or IP address)".
      * Haz clic en **Open** y se te pedirá el nombre de usuario y la contraseña en la ventana de terminal.

#### Proceso de Conexión

Una vez que ejecutes el comando `ssh`, el sistema remoto te pedirá tu contraseña:

```text
tu_usuario@192.168.1.50's password: **[Escribe tu contraseña]**
```

Tras ingresar la contraseña correctamente, habrás iniciado sesión.

-----

## Actualización del Sistema

Estos comandos son cruciales para mantener tu servidor **seguro y actualizado**.

| Comando | Función | Descripción |
| :--- | :--- | :--- |
| `sudo apt update` | **Actualiza la lista** | Descarga la información más reciente de los paquetes desde los repositorios. |
| `sudo apt upgrade` | **Actualiza paquetes** | Instala las versiones más nuevas de los paquetes que ya tienes instalados. |
| `sudo apt autoremove` | **Limpia dependencias** | Elimina automáticamente los paquetes que se instalaron como dependencias pero que ya no son necesarios. |

> **Importante:** El comando `sudo` se utiliza para ejecutar comandos con **privilegios de superusuario** (administrador).

-----

Este contenido ya fue cubierto en la sección **Navegación de Directorios** y la sección **Edición de Archivos con Nano**.

Para evitar redundancia, he fusionado la información de navegación existente con los comandos básicos de **administración de archivos y directorios** (crear, copiar, mover, eliminar), y he renombrado la sección para reflejar su contenido completo.

Aquí tienes la sección actualizada:

---

## Navegación y Administración de Archivos

Estos comandos te permiten moverte, ver el contenido dentro del sistema de archivos, y manipular archivos y directorios.

| Comando | Función | Ejemplo de Uso | Descripción |
| :--- | :--- | :--- | :--- |
| `pwd` | **Directorio actual** | `pwd` | Muestra la **ruta absoluta** del directorio donde te encuentras. |
| `ls` | **Listar contenido** | `ls -lha` | Muestra archivos y carpetas. Argumentos comunes: `-l` (largo), `-h` (tamaños legibles), `-a` (incluye ocultos). |
| `cd` | **Cambiar directorio** | `cd /var/log` | Se utiliza para **navegar** a otro directorio. |
| `cd ..` | **Subir un nivel** | `cd ..` | Mueve al **directorio padre** (superior). |
| `cd ~` | **Directorio Home** | `cd ~` | Te lleva rápidamente al **directorio personal** de tu usuario. |
| `cd -` | **Directorio anterior** | `cd -` | Vuelve al **último directorio** visitado. |

### Comandos de Manipulación de Archivos

| Comando | Función | Ejemplo de Uso | Descripción |
| :--- | :--- | :--- | :--- |
| `mkdir` | **Crear Directorio** | `mkdir nueva_carpeta` | Crea un nuevo directorio (carpeta). |
| `touch` | **Crear Archivo** | `touch nuevo_archivo.txt` | Crea un archivo vacío o actualiza la marca de tiempo de uno existente. |
| `cp` | **Copiar** | `cp archivo.txt /tmp/` | Copia un archivo a una nueva ubicación. Se usa `cp -r` para copiar directorios completos (recursivo). |
| `mv` | **Mover o Renombrar** | `mv archivo.txt nuevo_nombre.txt` | Renombra un archivo o directorio. |
| `mv` | **Mover o Renombrar** | `mv archivo.txt /home/usuario/docs` | Mueve un archivo a una ubicación diferente. |
| `rm` | **Eliminar Archivo** | `rm archivo_viejo.txt` | Elimina un archivo. |
| `rm -r` | **Eliminar Directorio** | `rm -r carpeta_vieja` | **Elimina un directorio y todo su contenido** (comando peligroso, usar con precaución). |

#### Rutas Comunes

* **`/`**: El **directorio raíz**, el punto de partida de todo el sistema de archivos.
* **`~`**: Un alias para tu directorio *Home*.

---

### Edición de Archivos con Nano

**Nano** es el editor de texto más sencillo para la terminal. Es común usarlo con `sudo` para editar archivos del sistema (que requieren permisos de administrador).

#### Abrir un Archivo (con Permisos)

```bash
sudo nano /etc/nginx/nginx.conf
```

  * Reemplaza la ruta con la del archivo que deseas editar.

#### Comandos Esenciales de Nano

Los comandos de Nano se activan pulsando la tecla **Ctrl** y la letra indicada (mostrada en la parte inferior del editor, por ejemplo, `^X` significa Ctrl+X).

| Combinación | Función | Descripción |
| :--- | :--- | :--- |
| **Ctrl + O** | **Guardar (Write Out)** | Guarda los cambios realizados en el archivo sin salir del editor. |
| **Ctrl + X** | **Salir (Exit)** | Intenta salir de Nano. Si hay cambios sin guardar, Nano te preguntará si deseas guardarlos. |
| **Ctrl + K** | **Cortar línea** | Corta toda la línea donde se encuentra el cursor. |
| **Ctrl + U** | **Pegar** | Pega el texto cortado. |

**Proceso para salir:**

1.  Presiona **Ctrl + X**.
2.  Si hiciste cambios, el editor te preguntará: `Save modified buffer? (Y/N/C)`
      * Presiona **Y** (Yes) para guardar los cambios.
      * Presiona **N** (No) para salir sin guardar.
      * Presiona **C** (Cancel) para volver al editor.

-----

## Clonación de Repositorios Git

Clonar un repositorio de Git es crear una copia local completa de un repositorio remoto usando el comando `git clone`.

### 1\. Obtener la URL del Repositorio

La URL se encuentra en la página del proyecto (GitHub, GitLab, Bitbucket).

  * Busca el botón **"Code"** o **"Clone"**.
  * Opciones comunes:
      * **HTTPS:** `https://github.com/usuario/mi-repo.git` (Más fácil si no has configurado SSH).
      * **SSH:** `git@github.com:usuario/mi-repo.git` (Requiere autenticación con claves SSH).

### 2\. Ejecutar el Comando `git clone`

Abre tu terminal, navega hasta el directorio donde deseas guardar la copia, y ejecuta el comando.

#### Formato Básico

```bash
git clone https://aws.amazon.com/es/what-is/repo/
```

#### Ejemplo (Usando HTTPS)

```bash
git clone https://github.com/usuario/mi-repo.git
```

#### Opciones Adicionales Útiles

| Comando | Función | Ejemplo |
| :--- | :--- | :--- |
| `[nombre_local]` | **Renombrar Carpeta:** Clona la *repo* en una carpeta local con un nombre diferente al original. | `git clone ... mi-nuevo-nombre` |
| `--depth 1` | **Clonación Superficial:** Solo descarga el historial más reciente. Útil para repositorios grandes. | `git clone --depth 1 ...` |

### 3\. Moverte al Repositorio Local

Git creará una carpeta con el nombre del repositorio. Usa `cd` para ingresar y empezar a trabajar:

```bash
cd mi-repo
```

---

## Ejecución y Construcción de Programas Go

Para trabajar con código Go, se utilizan los siguientes comandos en la terminal. Se asume que estás dentro del directorio raíz de tu proyecto donde se encuentra el archivo `main.go`.

### 1\. Ejecutar el Programa (Sin Compilar)

Puedes ejecutar el código fuente directamente sin generar un binario ejecutable. Esto es útil para pruebas rápidas.

| Comando | Función |
| :--- | :--- |
| `go run main.go` | Compila y ejecuta el archivo `main.go` al instante. El binario generado es temporal y se elimina al finalizar. |
| `go run .` | Ejecuta el paquete principal (`main`) en el directorio actual. |

### 2\. Construir el Binario (Compilar)

Para crear un archivo ejecutable independiente que pueda ser distribuido y ejecutado sin necesidad de tener Go instalado, usas el comando `go build`.

| Comando | Función |
| :--- | :--- |
| `go build main.go` | Compila el archivo `main.go` y genera un archivo ejecutable llamado `main` (o `main.exe` en Windows) en el mismo directorio. |
| `go build .` | Compila el paquete en el directorio actual y genera un ejecutable con el nombre del directorio. |

> **Nota:** Después de construir el binario, lo ejecutas directamente como cualquier otro programa:
>
> ```bash
> ./main
> ```

---

## Configuración de IP Estática con Netplan

En Ubuntu Server (a partir de la versión 17.10), la configuración de red se gestiona a través de **Netplan**, que utiliza archivos YAML sencillos.

### 1\. Identificar la Interfaz de Red

Primero, identifica el nombre de tu interfaz de red. Comúnmente es algo como `eth0`, `ensX`, o `enpXsY`.

```bash
ip a
```

### 2\. Localizar y Editar el Archivo de Configuración

Los archivos de configuración de Netplan se encuentran en `/etc/netplan/`.

#### A. Localizar el Nombre del Archivo

Para saber cómo se llama el archivo de configuración en tu sistema, lista el contenido del directorio. Generalmente sigue un formato numerado (ej. `00-installer-config.yaml`):

```bash
ls /etc/netplan/
```

#### B. Editar el Archivo

Utiliza `sudo nano` para editar el archivo exacto con permisos de administrador:

```bash
sudo nano /etc/netplan/nombre_del_archivo.yaml
```

### 3\. Asignar los Parámetros

Dentro del archivo, modifica la configuración de tu interfaz (reemplaza `ens33` con el nombre de tu interfaz y los valores con los de tu red). La **indentación** (espacios) es crítica en YAML.

```yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    ens33:                  # Nombre de tu interfaz de red
      dhcp4: no             # Deshabilita DHCP para usar IP estática
      addresses:
        - 192.168.1.100/24  # Dirección IP FIJA / Máscara de Red (CIDR)
      routes:
        - to: default
          via: 192.168.1.1  # Dirección IP del Gateway/Router
      nameservers:
        addresses:
          - 8.8.8.8         # DNS Principal
          - 8.8.4.4         # DNS Secundario
```

| Parámetro | Descripción |
| :--- | :--- |
| `addresses` | **IP Fija y Máscara de Red** (en formato CIDR). Por ejemplo, `/24` equivale a la máscara `255.255.255.0`. |
| `via` | **Gateway** (Puerta de Enlace) o dirección IP de tu *router*. |
| `nameservers: addresses` | **DNS Principal y Secundario**. Se usan los servidores de Google como ejemplo. |

### 4\. Aplicar los Cambios

Guarda el archivo (**Ctrl + O**, luego **Enter**) y sal de Nano (**Ctrl + X**). Luego, aplica la nueva configuración:

```bash
sudo netplan apply
```

Si la sintaxis YAML es correcta, la nueva configuración de red estática se aplicará inmediatamente.

---

¡Claro\! Cuando trabajas en un servidor de producción, el proceso de actualización de una webapp Go que usa Systemd para su servicio es de tres pasos clave: **construir el nuevo binario**, **detener el servicio**, y **reiniciar el servicio**.

Aquí tienes el manual que describe este flujo de trabajo.

-----

## Ciclo de Actualización y Despliegue de una Webapp Go (Systemd)

### 1\. Construir el Binario Actualizado

Debes regenerar el archivo ejecutable (`binario`) de tu aplicación para incluir los últimos cambios del código fuente.

**A. Navega al Directorio del Código Fuente:**

```bash
cd /ruta/a/tu/proyecto/go
```

**B. Genera el Nuevo Binario:**

Utiliza el comando `go build` para compilar el código. El flag `-o` se usa para especificar la ubicación y el nombre exacto del archivo ejecutable que Systemd está esperando.

```bash
go build -o /usr/local/bin/nombre_de_la_app main.go
```

| Comando | Descripción |
| :--- | :--- |
| `go build` | Compila el código Go. |
| `-o /usr/local/bin/nombre_de_la_app` | **Output:** Especifica la ruta y el nombre donde se guardará el binario compilado. Esta debe ser la misma ruta que se utiliza en el archivo de servicio de Systemd. |
| `main.go` | El archivo principal de tu aplicación. |

> **Nota:** La compilación suele ser muy rápida y si no hay errores, no verás ninguna salida.

### 2\. Reiniciar el Servicio con Systemd

Una vez que el nuevo binario está en su lugar, debes decirle a Systemd que detenga la versión antigua de la aplicación y que inicie la nueva.

| Comando | Función |
| :--- | :--- |
| `sudo systemctl stop nombre_del_servicio.service` | **Detener:** Detiene inmediatamente la ejecución de la versión anterior de la aplicación. |
| `sudo systemctl start nombre_del_servicio.service` | **Iniciar:** Inicia la aplicación usando el nuevo binario que acabas de construir. |
| `sudo systemctl restart nombre_del_servicio.service` | **Reinicio Único:** La opción recomendada. Detiene y luego inicia el servicio en un solo paso. |

#### Paso Recomendado (Reiniciar):

```bash
sudo systemctl restart nombre_del_servicio.service
```

### 3\. Verificar el Estado del Servicio y Logs

Después del reinicio, es vital asegurarse de que la aplicación haya comenzado correctamente y no tenga errores.

**A. Verificar el Estado:**

Comprueba si el servicio está en estado `active (running)`.

```bash
sudo systemctl status nombre_del_servicio.service
```

**B. Revisar los Logs (Registro):**

Si la aplicación no se inicia, revisa los mensajes de error en los logs del sistema (`journalctl`). Esto es crucial para la depuración.

```bash
sudo journalctl -u nombre_del_servicio.service --since "5 minutes ago"
```

> Reemplaza `nombre_del_servicio.service` (ej. `mi-webapp.service`) por el nombre real de tu archivo de unidad Systemd.

-----

Aquí tienes la sección que describe cómo verificar la accesibilidad de tu servidor y tu aplicación web desde una máquina remota, utilizando comandos básicos de red.

-----

## Verificación de Conectividad y Webapp

Una vez que el servidor tiene la IP fija y la aplicación está desplegada, debes verificar su estado desde una computadora remota.

### 1\. Verificar el Estado del Servidor (Ping)

El comando `ping` verifica la **conectividad básica** (Capa 3 - Red) al servidor. Si este comando falla, ningún otro servicio será accesible.

| Sistema | Comando |
| :--- | :--- |
| **Windows, Linux, macOS** | `ping [Dirección IP del Servidor]` |

**Ejemplo:**

```bash
ping 192.168.1.100
```

  * **Si es exitoso:** Verás respuestas (`reply`) con el tiempo de latencia. Esto confirma que el servidor está encendido, la red funciona, y el **firewall no está bloqueando ICMP** (el protocolo que usa `ping`).
  * **Si falla:** El servidor está apagado, la red está mal configurada, o el firewall (`ufw`) está bloqueando el tráfico `ping`.

### 2\. Verificar la Webapp y los Puertos (Telnet o netcat)

El `ping` solo prueba que el servidor existe; no prueba que tu aplicación esté funcionando. Para eso, debes verificar que el puerto que usa tu webapp (comúnmente **80** para HTTP o **443** para HTTPS, o un puerto personalizado como **8080**) esté abierto.

#### Opción A: Usar Telnet

Telnet intenta establecer una conexión al puerto especificado. Si la conexión se establece (la pantalla se queda en blanco o muestra un mensaje de bienvenida), el puerto está abierto.

| Sistema | Comando |
| :--- | :--- |
| **Linux/macOS** | `telnet [Dirección IP del Servidor] [Puerto]` |

**Ejemplo (Webapp en el puerto 8080):**

```bash
telnet 192.168.1.100 8080
```

#### Opción B: Usar netcat (nc)

`netcat` es una alternativa más moderna y flexible a Telnet.

| Sistema | Comando |
| :--- | :--- |
| **Linux/macOS** | `nc -vz [Dirección IP del Servidor] [Puerto]` |

**Ejemplo:**

```bash
nc -vz 192.168.1.100 8080
```

  * **Si es exitoso:** La salida indicará `Connection succeeded` o el terminal se conectará. Esto confirma que tu **aplicación está corriendo** y el **firewall está permitiendo el acceso** a ese puerto específico.

### 3\. Verificar en el Navegador

Finalmente, verifica la accesibilidad de la aplicación web abriendo tu navegador:

```
http://[Dirección IP del Servidor]:[Puerto]
```

**Ejemplo:**

```
http://192.168.1.100:8080
```

Si ves la interfaz de tu aplicación, el despliegue ha sido exitoso.

---

Aquí tienes una sección concisa sobre cómo cambiar entre diferentes ventanas o consolas (TTYs) en Ubuntu Server cuando accedes directamente a la máquina.

---

## Cambio entre Terminales (TTYs) en Consola Local

Cuando trabajas directamente en el servidor (sin SSH), puedes acceder a **seis terminales virtuales** o TTYs (Teletipos) diferentes. Esto es útil si necesitas iniciar sesión con otro usuario o realizar una tarea sin interrumpir una aplicación que se está ejecutando en la consola actual.

### Cómo Cambiar de Ventana

Para cambiar entre estas terminales virtuales, utiliza la combinación de teclas **Ctrl + Alt + Función (Fn)**, donde el número de la tecla de función (`F1` a `F6`) corresponde al número de la terminal que deseas abrir.

| Tecla | Función | Descripción |
| :--- | :--- | :--- |
| **Ctrl + Alt + F1** | **TTY1 (Principal)** | A menudo muestra información del *boot* o un *login prompt* (inicio de sesión). |
| **Ctrl + Alt + F2** | **TTY2** | La consola de trabajo más común después del *boot*. Aquí es donde típicamente inicias sesión. |
| **Ctrl + Alt + F3** | **TTY3** | Otra consola de inicio de sesión disponible. |
| **Ctrl + Alt + F4** | **TTY4** | Otra consola disponible. |
| **Ctrl + Alt + F5** | **TTY5** | Otra consola disponible. |
| **Ctrl + Alt + F6** | **TTY6** | La última consola de inicio de sesión disponible. |

### Consideraciones

* **Entorno Gráfico:** Si en algún momento instalas una interfaz gráfica (GUI) en Ubuntu Server, esta generalmente ocupará la consola **TTY7** (accesible con `Ctrl + Alt + F7`). El servidor puro usa de F1 a F6.
* **Sesiones Independientes:** Cada TTY funciona como una **sesión completamente independiente**. Si inicias sesión en TTY2 y luego cambias a TTY3, deberás iniciar sesión nuevamente en TTY3.
* **TTY Actual:** La terminal que está mostrando tu *prompt* o la pantalla de inicio de sesión es tu TTY activa. Solo puedes tener una TTY activa a la vez.

---

Aquí tienes la sección sobre cómo agregar nuevos repositorios e instalar software utilizando el gestor de paquetes **APT** en Ubuntu Server.

-----

## Gestión de Paquetes y Repositorios con APT

En Ubuntu, la instalación de software se realiza principalmente a través de la herramienta **APT (Advanced Package Tool)**, que interactúa con repositorios de software.

### 1\. Agregar Nuevos Repositorios

A veces, el software que necesitas no está en los repositorios predeterminados de Ubuntu. Para añadir repositorios externos (a menudo llamados **PPA - Personal Package Archive** o repositorios de terceros), se utiliza el comando `add-apt-repository`.

#### A. Instalar Herramientas (si es necesario)

Si es la primera vez que agregas un PPA, es posible que necesites instalar una herramienta auxiliar:

```bash
sudo apt update
sudo apt install software-properties-common
```

#### B. Añadir el Repositorio

Usa el siguiente formato, reemplazando la PPA por la que necesitas:

```bash
sudo add-apt-repository ppa:[nombre_del_ppa/ppa]
```

**Ejemplo:** Para añadir un repositorio para una versión específica de un programa:

```bash
sudo add-apt-repository ppa:ondrej/php
```

### 2\. Sincronizar y Actualizar

Después de añadir cualquier repositorio nuevo (o si ha pasado tiempo desde la última vez), debes **sincronizar** tu lista local de paquetes con los repositorios remotos:

```bash
sudo apt update
```

> **Nota:** Este comando **solo descarga la información** de los paquetes (qué versiones están disponibles), pero **no instala** software nuevo.

### 3\. Instalación de Programas

Una vez que los repositorios están actualizados, utiliza `apt install` para descargar e instalar el software.

| Comando | Función |
| :--- | :--- |
| `sudo apt install [nombre_del_paquete]` | **Instalar:** Descarga e instala un paquete específico, incluyendo todas sus dependencias. |
| `sudo apt install -y [paquete]` | **Instalar sin Confirmar:** Instala el paquete automáticamente sin preguntar la confirmación `[Y/n]`. |

**Ejemplo:** Instalar el servidor web Nginx:

```bash
sudo apt install nginx
```

### 4\. Eliminar Programas

Para desinstalar software que ya no necesitas:

| Comando | Función |
| :--- | :--- |
| `sudo apt remove [nombre_del_paquete]` | **Desinstalar:** Elimina los archivos binarios del paquete, pero **mantiene** los archivos de configuración. |
| `sudo apt purge [nombre_del_paquete]` | **Eliminar Completamente:** Elimina el paquete y también sus **archivos de configuración**. Recomendado para una limpieza total. |
| `sudo apt autoremove` | **Limpieza de Dependencias:** Elimina automáticamente paquetes que se instalaron como dependencias y que ya no son requeridos por ningún otro programa instalado. |

---

Aquí tienes una nueva sección que explica cómo acceder a la consola de MySQL/MariaDB en Ubuntu Server para ejecutar *queries* y cómo habilitar el acceso remoto para conectarte usando MySQL Workbench.

-----

## Administración de Bases de Datos MySQL

### 1\. Acceso a la Consola Local

Para administrar y ejecutar comandos SQL en tu servidor Ubuntu, accede a la consola de MySQL/MariaDB desde la terminal del servidor:

| Comando | Función | Descripción |
| :--- | :--- | :--- |
| `sudo mysql` | **Root Local:** Abre la consola como usuario *root* del sistema (no requiere contraseña de MySQL/MariaDB). |
| `mysql -u usuario -p` | **Usuario Específico:** Abre la consola solicitando la contraseña del usuario de base de datos especificado. |

Una vez dentro de la consola, puedes realizar las siguientes acciones:

| Comando SQL | Función |
| :--- | :--- |
| `SHOW DATABASES;` | Muestra todas las bases de datos en el servidor. |
| `USE nombre_db;` | Selecciona la base de datos con la que deseas trabajar. |
| `SHOW TABLES;` | Lista todas las tablas dentro de la base de datos seleccionada. |
| `SELECT * FROM tabla;` | Realiza una *query* de ejemplo. |
| `EXIT;` | Cierra la consola de MySQL/MariaDB y regresa a la terminal de Ubuntu. |

### 2\. Habilitar el Acceso Remoto (MySQL Workbench)

Para conectar herramientas gráficas externas como MySQL Workbench, debes hacer dos cosas: **Configurar el usuario** en la base de datos y **configurar el servidor** para que escuche conexiones externas.

#### A. Configurar Usuario para Conexiones Remotas

Por defecto, los usuarios de MySQL suelen estar limitados a `localhost`. Necesitas crear o modificar un usuario para que acepte conexiones desde cualquier IP (`%`).

1.  **Accede a la consola** de MySQL (`sudo mysql`).

2.  **Crea un usuario** remoto (reemplaza `usuario_remoto`, `password` y `nombre_db`):

    ```sql
    CREATE USER 'usuario_remoto'@'%' IDENTIFIED BY 'tu_contraseña_segura';
    GRANT ALL PRIVILEGES ON nombre_db.* TO 'usuario_remoto'@'%' WITH GRANT OPTION;
    FLUSH PRIVILEGES;
    EXIT;
    ```

#### B. Configurar el Archivo de Servicio MySQL

Por defecto, MySQL solo escucha conexiones de `localhost`. Debes editar el archivo de configuración para cambiar la directiva `bind-address`.

1.  **Edita el archivo de configuración:**

    ```bash
    sudo nano /etc/mysql/mysql.conf.d/mysqld.cnf
    ```

    (La ruta puede variar ligeramente, verifica `/etc/mysql/my.cnf` si no encuentras el archivo).

2.  **Modifica `bind-address`:**
    Busca la línea:

    ```ini
    bind-address            = 127.0.0.1
    ```

    Y cámbiala para que el servidor escuche todas las interfaces:

    ```ini
    bind-address            = 0.0.0.0
    ```

3.  **Reinicia el servicio MySQL:**

    ```bash
    sudo systemctl restart mysql
    ```

#### C. Abrir el Puerto en el Firewall (UFW)

Asegúrate de que tu firewall (`ufw`) permita las conexiones entrantes al puerto de MySQL (por defecto, **3306**).

```bash
sudo ufw allow 3306/tcp
sudo ufw reload
```

### 3\. Conexión desde MySQL Workbench

En tu computadora remota, abre MySQL Workbench y crea una nueva conexión:

  * **Hostname:** Dirección IP del Servidor Ubuntu (ej. `192.168.1.100`)
  * **Port:** `3306`
  * **Username:** El usuario remoto que creaste (ej. `usuario_remoto`)
  * **Password:** La contraseña configurada.
