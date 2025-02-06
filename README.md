<p style="text-align:center;" align="center"><a href="https://meshery.io"><picture align="center">
  <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-dark-text-side.svg"  width="70%" align="center" style="margin-bottom:20px;">
  <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-light-text-side.svg" width="70%" align="center" style="margin-bottom:20px;">
  <img alt="Shows an illustrated light mode meshery logo in light color mode and a dark mode meshery logo dark color mode." src="https://raw.githubusercontent.com/meshery/meshery/master/.github/assets/images/meshery/meshery-logo-tag-light-text-side.png" width="70%" align="center" style="margin-bottom:20px;">
</picture></a><br /><br /></p>
<p align="center">
 
<div align="center">

[![Docker Pulls](https://img.shields.io/docker/pulls/layer5/meshery-cilium.svg)](https://hub.docker.com/r/layer5/meshery-cilium)
[![Go Report Card](https://goreportcard.com/badge/github.com/layer5io/meshery-cilium)](https://goreportcard.com/report/github.com/meshery/meshery-cilium)
[![Build Status](https://img.shields.io/github/actions/workflow/status/meshery/meshery-cilium/multi-platform.yml?branch=master)](https://github.com/meshery/meshery-cilium/actions)
[![GitHub](https://img.shields.io/github/license/layer5io/meshery-istio.svg)](LICENSE)
[![GitHub issues by-label](https://img.shields.io/github/issues/layer5io/meshery-cilium/help%20wanted.svg)](https://github.com/meshery/meshery-cilium/issues?q=is%3Aopen+is%3Aissue+label%3A"help+wanted")
[![Website](https://img.shields.io/website/https/layer5.io/meshery.svg)](https://meshery.io)
[![Twitter Follow](https://img.shields.io/twitter/follow/layer5.svg?label=Follow&style=social)](https://twitter.com/intent/follow?screen_name=mesheryio)
[![Discuss Users](https://img.shields.io/discourse/users?server=https%3A%2F%2Fdiscuss.layer5.io)](https://meshery.io/community#discussion-forums)
[![Slack](https://img.shields.io/badge/Slack-@layer5.svg?logo=slack)](http://slack.meshery.io)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/3564/badge)](https://bestpractices.coreinfrastructure.org/projects/3564)

</div>

# Meshery Adapter for Cilium Service Mesh

<h2><a href="https://docs.meshery.io/">Cilium Service Mesh</h2>
<a href="https://cilium.io/">
 <img src="https://raw.githubusercontent.com/meshery/meshery-cilium/master/.github/readme/images/cilium.svg" style="margin:10px;" width="125px" 
   alt="Cilium logo" align="left" />
</a>
<p>
<a href="https://cilium.io/">Cilium</a> is open source software for providing, securing, and observing network connectivity between container workloads - cloud-native, and fueled by eBPF as Linux kernel technology. Cilium is bringing eBPF strengths to the world of service mesh. Cilium Service Mesh features eBPF-powered connectivity, traffic management, security, and observability.
 <br /> <br /> 
</p>

<p style="clear:both;">
<h2><a href="https://meshery.io/">Meshery</a></h2>
<a href="https://meshery.io"><img src="https://raw.githubusercontent.com/meshery/meshery-cilium/master/.github/readme/images/meshery/meshery-logo-light.svg"
style="margin:10px;" width="125px" 
alt="Meshery - the Cloud Native Management Plane" align="left" /></a>
A self-service engineering platform, <a href="https://meshery.io">Meshery</a>, is the open source, cloud native manager that enables the design and management of all Kubernetes-based infrastructure and applications (multi-cloud). Among other features, As an extensible platform, Meshery offers visual and collaborative GitOps, freeing you from the chains of YAML while managing Kubernetes multi-cluster deployments.
<br /><br /><p align="center"><i>If you’re using Meshery or if you like the project, please <a href="https://github.com/meshery/meshery/stargazers">★</a> star this repository to show your support! 🤩</i></p>
</p>
<br /><br />

## Join the Community!

<a name="contributing"></a><a name="community"></a>
Our projects are community-built and welcome collaboration. 👍 Be sure to see the <a href="https://docs.meshery.io/project/contributing#not-sure-where-to-start">Contributor Welcome Guide</a> and jump into our <a href="http://slack.meshery.io">Slack</a>!

<p style="clear:both;">
<a href="https://meshery.io/community"><img alt="MeshMates" src=".github\readme\images\Layer5-MeshMentors-1.png" style="margin-right:10px; margin-bottom:15px;" width="28%" align="left"/></a>
<h3>Find your MeshMate</h3>

<p>MeshMates are experienced community members, who will help you learn your way around, discover live projects and expand your community network. 
Become a <b>Meshtee</b> today!</p>

Find out more on the <a href="https://meshery.io/community">Meshery community</a>. <br />
<br /><br />

</p>

<div>&nbsp;</div>


<a href="https://slack.meshery.io">

<picture align="right">
  <source media="(prefers-color-scheme: dark)" srcset=".github\readme\images\slack-dark-128.png"  width="110px" align="right" style="margin-left:10px;margin-top:10px;">
  <source media="(prefers-color-scheme: light)" srcset=".github\readme\images\slack-128.png" width="110px" align="right" style="margin-left:10px;padding-top:5px;">
  <img alt="Shows an illustrated light mode meshery logo in light color mode and a dark mode meshery logo dark color mode." src=".github\readme\images\slack-128.png" width="110px" align="right" style="margin-left:10px;padding-top:13px;">
</picture>
</a>

<a href="https://meshery.io/community"><img alt="Meshery Cloud Native Community" src="https://raw.githubusercontent.com/meshery/meshery-cilium/master/.github/readme/images//community.svg" style="margin-right:8px;padding-top:5px;" width="140px" align="left" /></a>
<p>
✔️ <em><strong>Join</strong></em> any or all of the weekly meetings on <a href="https://meshery.io/calendar">community calendar</a>.<br />
✔️ <em><strong>Watch</strong></em> community <a href="https://www.youtube.com/@mesheryio?sub_confirmation=1">meeting recordings</a>.<br />
✔️ <em><strong>Fill-in</strong></em> a <a href="https://layer5.io/newcomers">community member form</a> to gain access to community resources.
<br />
✔️ <em><strong>Discuss</strong></em> in the <a href="https://meshery.io/community#discussion-forums">Community Forum</a>.<br />
✔️ <em><strong>Explore more</strong></em> in the <a href="https://meshery.io/community#handbook">Community Handbook</a>.<br />
</p>
<p align="center">
<i>Not sure where to start?</i> Grab an open issue with the <a href="https://github.com/issues?q=is%3Aopen+is%3Aissue+archived%3Afalse+org%3Alayer5io+org%3Ameshery+org%3Aservice-mesh-performance+org%3Aservice-mesh-patterns+label%3A%22help+wanted%22+">help-wanted label</a>.</p>

**License**

This repository and site are available as open source under the terms of the [Apache 2.0 License](https://opensource.org/licenses/Apache-2.0).
