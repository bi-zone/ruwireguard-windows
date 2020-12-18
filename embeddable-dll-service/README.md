## Embeddable WireGuard Tunnel Library

This allows embedding WireGuard as a service inside of another application. Build `tunnel.dll` by running `./build.bat` in this folder. The first time you run it, it will invoke `..\build.bat` simply for downloading dependencies. After, you should have `amd64/tunnel.dll` and `x86/tunnel.dll`.

The basic setup to use `tunnel.dll` is:

##### 1. Install a service with these parameters:

    Service Name:  "SomeServiceName"
    Display Name:  "Some Service Name"
    Service Type:  SERVICE_WIN32_OWN_PROCESS
    Start Type:    StartAutomatic
    Error Control: ErrorNormal,
    Dependencies:  [ "Nsi" ]
    Sid Type:      SERVICE_SID_TYPE_UNRESTRICTED
    Executable:    "C:\path\to\example\vpnclient.exe /service configfile.conf"

Some of these may have to be changed with `ChangeServiceConfig2` after the
initial call to `CreateService` The `SERVICE_SID_TYPE_UNRESTRICTED` parameter
is absolutely essential; do not forget it.

##### 2. Have your program's main function handle the `/service` switch:

    if (!strcmp(argv[1], "/service") && argc == 3) {
        HMODULE tunnel_lib = LoadLibrary("tunnel.dll");
        if (!tunnel_lib)
            abort();
        tunnel_proc_t tunnel_proc = (tunnel_proc_t)GetProcAddress(tunnel_lib, "WireGuardTunnelService");
        if (!tunnel_proc)
            abort();
        struct go_string conf_file = {
            .str = argv[2],
            .n = strlen(argv[2])
        };
        return tunnel_proc(conf_file);
    }

##### 3. Scoop up logs by implementing a ringlogger format reader.

##### 4. Talk to the service over its named pipe.

There is a sample implementation of bits and pieces of this inside of the `csharp\` directory.
