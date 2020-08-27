package compute

import (
	"testing"
)

func Test_computeArch(t *testing.T) {
	_, err := computeArch(Input{Arch: "xpto-invalid-arch"}, ExportCfg{})

	if err == nil {
		t.Error("should support only x86_64 and i686")
	}

	in := fakeInput()
	in.Arch = "i686"
	in.TotalRAM = 1 * TB

	out, _ := computeArch(*in, *NewExportCfg(*in))

	if out.Memory.SharedBuffers > 4*GB ||
		out.Memory.WorkMem > 4*GB ||
		out.Memory.MaintenanceWorkMem > 4*GB {
		t.Error("should limit ANY memory parameter to a max of 4GB in a 32 bits system")
	}

}