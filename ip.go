package main

import (
	"errors"
	"fmt"
	"math"
	"net"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type uiIP struct {
	app.Compo
	cidrInput   string
	broadCast   string
	networkSize int
}

func newUiIp() *uiIP {
	return &uiIP{}
}
func (h *uiIP) Render() app.UI {
	return app.Div().Class("container").Styles(map[string]string{"width": "100%", "max-width": "400px", "padding": "15px", "margin": "auto"}).Body(
		app.Form().Body(
			app.H1().Class("h3 mb-3 fw-normal text-center").Body(app.Text("IP Tool")),
			app.Div().Class("form-floating mb-3").Body(
				app.Input().ID("cidr").Type("text").OnChange(h.ValueTo(&h.cidrInput)).Value(h.cidrInput).Class("form-control"),
				app.Label().Body(app.Text("Enter CIDR:")).For("cidr"),
			),
			app.Button().Class("w-100 btn btn-lg btn-primary").Type("Submit").OnClick(h.eventCidrHandler).Body(app.Text("Submit")),
			app.Div().Class("form-floating mb-3 mt-3").Body(
				app.Input().ID("broadcast").Type("text").ReadOnly(true).Value(h.broadCast).Class("form-control"),
				app.Label().Body(app.Text("Broadcast:")).For("broadcast"),
			),
			app.Div().Class("form-floating mb-3").Body(
				app.Input().ID("size").Type("text").ReadOnly(true).Value(h.networkSize).Class("form-control"),
				app.Label().Body(app.Text("Size")).For("size"),
			),
		).OnSubmit(h.onSubmit),
		getFooter(),
	)
}

func (h *uiIP) OnKeyPress(ctx app.Context, e app.Event) {
	code := e.Get("key").String()
	fmt.Printf("%q", code)
	if code == "Enter" {
		fmt.Print("HIHIHI")
		h.eventCidrHandler(ctx, e)
	}
}

func (h *uiIP) eventCidrHandler(ctx app.Context, e app.Event) {
	var err error
	_, h.broadCast, h.networkSize, err = calculateSubnetDetails(h.cidrInput)
	if err != nil {
		fmt.Println(h.cidrInput)
	}
	ctx.Update()

}

func (h *uiIP) onSubmit(ctx app.Context, e app.Event) {
	e.PreventDefault()
	h.eventCidrHandler(ctx, e)
}

// calculateSubnetDetails takes a CIDR string and returns the subnet IP, broadcast IP, and subnet size.
func calculateSubnetDetails(cidr string) (subnetIP net.IP, broadcastIP string, size int, err error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, "", 0, fmt.Errorf("error parsing CIDR: %v", err)
	}

	subnetIP = ip
	bIP, err := getBroadcastAddress(ipnet)
	if err != nil {
		return subnetIP, broadcastIP, size, err
	}
	size, err = getNetworkSize(ipnet)

	return subnetIP, bIP.String(), size, err // Subtract 2 for the network and broadcast addresses
}

// getBroadcastAddress takes a net.IPNet and returns the broadcast address as net.IP.
func getBroadcastAddress(ipNet *net.IPNet) (net.IP, error) {
	if ipNet == nil {
		return nil, errors.New("input IPNet cannot be nil")
	}

	ip := ipNet.IP.To4()
	if ip == nil {
		return nil, errors.New("only IPv4 is supported")
	}

	mask := ipNet.Mask
	broadcast := make(net.IP, len(ip))

	// Calculate the broadcast address by OR'ing the IP with the inverse of the subnet mask
	for i := 0; i < len(ip); i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}

	return broadcast, nil
}

// getNetworkSize returns the number of usable host addresses within a given net.IPNet.
func getNetworkSize(ipNet *net.IPNet) (int, error) {
	if ipNet == nil {
		return 0, errors.New("input IPNet cannot be nil")
	}

	ip := ipNet.IP.To4()
	if ip == nil {
		return 0, errors.New("only IPv4 is supported")
	}

	// Calculate the number of hosts by inverting the subnet mask and adding 1
	ones, bits := ipNet.Mask.Size()
	if bits != 32 {
		return 0, errors.New("incorrect IP length for IPv4")
	}

	totalAddresses := int(math.Pow(2, float64(bits-ones)))

	// Subtract 2 to account for the network and broadcast addresses, if applicable
	if totalAddresses <= 2 {
		return 0, nil // No usable addresses in /31 or /32 for traditional networks
	}

	return totalAddresses - 2, nil
}
