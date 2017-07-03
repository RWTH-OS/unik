# uHyve Provider
UniK supports running HermitCore unikernels through uHyve.
uHyve is a thin hypervisor powered by Linux's KVM API.

To run UniK instances with uHyve, change `daemon-config.yaml`:

```yaml
providers:
  #...
  uhyve:
    - name: uhyve-name
```

To run:

```
unik build --name hermitcore_example --path ./csrc --base hermitcore --language c --provider uhyve
```
