package factory

func (f *Factory) GetServiceRequest() *ServiceRequest {
	return &f.sr
}

func (f *Factory) GetConfiguration() *Configuration {
	return f.sr.Configuration
}

func (f *Factory) GetGslbService() *GSLBService {
	if f.GetConfiguration() == nil {
		return nil
	}
	return f.sr.Configuration.GslbService
}
