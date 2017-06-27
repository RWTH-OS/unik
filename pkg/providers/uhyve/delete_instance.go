package uhyve

func (p *UhyveProvider) DeleteInstance(id string, force bool) error {
	return p.StopInstance(id)
}
